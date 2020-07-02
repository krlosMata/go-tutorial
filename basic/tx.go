package basic

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Tx struct {
	From   common.Address
	To     common.Address
	Amount uint64
	Hash   common.Hash
}

func NewTx(from common.Address, to common.Address, amount uint64) Tx {

	tx := Tx{
		from,
		to,
		amount,
		common.Hash{},
	}

	b := tx.Bytes()

	tx.Hash = crypto.Keccak256Hash(b[:])

	return tx
}

func NewTxFromBytes(b []byte) Tx {
	arrFrom := common.Address{}
	copy(arrFrom[:], b[0:20])

	arrTo := common.Address{}
	copy(arrTo[:], b[20:40])

	amount := binary.LittleEndian.Uint64(b[40:48])

	hash := common.Hash{}
	copy(hash[:], b[48:80])

	return Tx{
		arrFrom,
		arrTo,
		amount,
		hash,
	}
}

func (tx *Tx) Bytes() []byte {
	var b []byte
	b = append(b, tx.From[:]...)
	b = append(b, tx.To[:]...)

	bytesAmount := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesAmount, tx.Amount)

	b = append(b, bytesAmount[:]...)
	b = append(b, tx.Hash[:]...)

	return b
}

func (tx *Tx) Hex() string {
	return hex.EncodeToString(tx.Bytes())
}
