package basic

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
)

const keyTx string = "tx-"

// Service store data on database
type Service struct {
	Db leveldb.DB
}

// NewService start database
func NewService(pathDb string) (*Service, error) {
	db, err := leveldb.OpenFile(pathDb, nil)
	if err != nil {
		return nil, err
	}

	s := Service{*db}

	return &s, nil
}

// UpdateBalances add tokens to a specified address
func (s *Service) UpdateBalances(address common.Address, amount uint64) error {

	balance, err := s.GetBalance(address)
	if err != nil {
		return err
	}

	balanceFinal := balance + amount

	bytesBalance := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesBalance, balanceFinal)

	err = s.Db.Put([]byte(address.String()), bytesBalance, nil)

	if err != nil {
		return err
	}

	return nil
}

// AddTx Verify and add tx.  Updates balances.
func (s *Service) AddTx(tx Tx) error {
	isVerified := tx.VerifyTx()
	if isVerified == false {
		return fmt.Errorf("Signature is not correct")
	}

	from := tx.From
	to := tx.To
	amount := tx.Amount

	// check enough balance from sender account
	balanceFrom, err := s.GetBalance(from)
	if err != nil {
		return err
	}

	checkBalance := int64(balanceFrom) - int64(amount)

	if checkBalance < 0 {
		return fmt.Errorf("Not enough funds on sender address")
	}

	balanceTo, err := s.GetBalance(to)
	if err != nil {
		return err
	}

	balanceFinalFrom := balanceFrom - amount
	balanceFinalTo := balanceTo + amount

	s.UpdateBalances(from, balanceFinalFrom)
	s.UpdateBalances(to, balanceFinalTo)

	// check enough balance from sender account
	txFrom, err := s.Db.Get([]byte(keyTx+from.String()), nil)

	if err != leveldb.ErrNotFound && err != nil {
		return err
	}

	var b []byte
	if err == leveldb.ErrNotFound {
		txFrom = []byte{}
	}

	b = append(b, txFrom[:]...)
	b = append(b, tx.Bytes()...)

	err = s.Db.Put([]byte(keyTx+from.String()), b, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetBalance for an address
func (s *Service) GetBalance(address common.Address) (uint64, error) {

	amountDb, err := s.Db.Get([]byte(address.String()), nil)

	// if address is not found, balance is 0
	if err == leveldb.ErrNotFound {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(amountDb), nil
}

// GetTxs all transactions from an specific address
func (s *Service) GetTxs(address common.Address) ([]Tx, error) {
	// check enough balance from sender account
	txs, err := s.Db.Get([]byte(keyTx+address.String()), nil)

	if err != leveldb.ErrNotFound && err != nil {
		return nil, err
	}

	// return empty array if no transaction is found
	if err == leveldb.ErrNotFound {
		return []Tx{}, nil
	}

	numTx := len(txs) / lengthTx

	var retTxs []Tx = []Tx{}

	for i := 0; i < numTx; i++ {
		tx := NewTxFromBytes(txs[i*lengthTx : (i+1)*lengthTx])
		retTxs = append(retTxs, tx)
	}

	return retTxs, nil
}
