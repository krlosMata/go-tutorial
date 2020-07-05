package basic

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Tx transsaction struct
type Tx struct {
	From      common.Address               `json:"from"`
	To        common.Address               `json:"to"`
	Amount    uint64                       `json:"amount"`
	Hash      common.Hash                  `json:"hash"`
	Signature [crypto.SignatureLength]byte `json:"signature"`
	Timestamp uint64                       `json:"timestamp"`
	Nonce     uint64                       `json:"nonce"`
}

// NewTx create transaction from basic fields
func NewTx(from common.Address, to common.Address, amount uint64) Tx {
	signature := [crypto.SignatureLength]byte{}
	timestamp := time.Now().Unix()

	tx := Tx{
		from,
		to,
		amount,
		common.Hash{},
		signature,
		uint64(timestamp),
		0,
	}

	// get nonce with first bytes at 0
	target := []byte{0}
	hash := common.Hash{}
	for {
		hash = tx.calcHash()
		res := bytes.Compare(hash[0:1], target)
		if res == 0 {
			break
		}
		tx.Nonce++
	}
	tx.Hash = hash

	return tx
}

// NewTxFromBytes loads transaction from bytes
func NewTxFromBytes(b []byte) Tx {
	arrFrom := common.Address{}
	copy(arrFrom[:], b[0:20])

	arrTo := common.Address{}
	copy(arrTo[:], b[20:40])

	amount := binary.LittleEndian.Uint64(b[40:48])

	hash := common.Hash{}
	copy(hash[:], b[48:80])

	arrSignature := [crypto.SignatureLength]byte{}
	copy(arrSignature[:], b[80:80+crypto.SignatureLength])

	timestamp := binary.LittleEndian.Uint64(b[80+crypto.SignatureLength : 80+crypto.SignatureLength+8])
	nonce := binary.LittleEndian.Uint64(b[80+crypto.SignatureLength+8 : 80+crypto.SignatureLength+8+8])

	return Tx{
		arrFrom,
		arrTo,
		amount,
		hash,
		arrSignature,
		timestamp,
		nonce,
	}
}

func (tx *Tx) calcHash() common.Hash {
	var b []byte
	// add From
	b = append(b, tx.From[:]...)
	// add To
	b = append(b, tx.To[:]...)

	// add Amount
	bytesAmount := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesAmount, tx.Amount)
	b = append(b, bytesAmount[:]...)

	// add Timestamp
	bytesTimestamp := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesTimestamp, tx.Timestamp)
	b = append(b, bytesTimestamp[:]...)

	// add Nonce
	bytesNonce := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesNonce, tx.Nonce)
	b = append(b, bytesNonce[:]...)

	return crypto.Keccak256Hash(b[:])
}

// Bytes get transaction in an array of bytes
func (tx *Tx) Bytes() []byte {
	var b []byte
	b = append(b, tx.From[:]...)
	b = append(b, tx.To[:]...)

	bytesAmount := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesAmount, tx.Amount)

	b = append(b, bytesAmount[:]...)
	b = append(b, tx.Hash[:]...)
	b = append(b, tx.Signature[:]...)

	bytesTimestamp := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesTimestamp, tx.Timestamp)
	b = append(b, bytesTimestamp[:]...)

	bytesNonce := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesNonce, tx.Nonce)
	b = append(b, bytesNonce[:]...)

	return b
}

// Hex trandaction to Hex string
func (tx *Tx) Hex() string {
	return hex.EncodeToString(tx.Bytes())
}

// VerifyTx check `from` address is the message signer
func (tx *Tx) VerifyTx() bool {
	publicKeyBytes, err := crypto.Ecrecover(tx.Hash[:], tx.Signature[:])
	if err != nil {
		return false
	}

	pub, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		return false
	}

	address := crypto.PubkeyToAddress(*pub)
	return address == tx.From
}
