package basic

import (
	"crypto/ecdsa"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Account struct
type Account struct {
	PrivateKey ecdsa.PrivateKey `json:"privateKey"`
	PublicKey  ecdsa.PublicKey  `json:"publicKey"`
	EthAddress common.Address   `json:"ethereumAddress"`
}

// NewAccountRandom create new account from random private key
func NewAccountRandom() Account {
	priv, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(priv.PublicKey)

	return Account{
		*priv,
		priv.PublicKey,
		address,
	}
}

// SignTx Sign transaction and adds the signature
func (a *Account) SignTx(tx *Tx) {
	hash := tx.Hash
	signature, err := crypto.Sign(hash[:], &a.PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	var arr [crypto.SignatureLength]byte
	copy(arr[:], signature[:])
	tx.Signature = arr
}
