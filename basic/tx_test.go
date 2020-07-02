package basic

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
)

func TestNewTx(t *testing.T) {
	from := common.HexToAddress("0xcc56f727B11289925941D9100b992b0e60ff1C8e")
	to := common.HexToAddress("0x40eF59C06031BD167Ed206F404f2A7f38b377Df9")
	amount := uint64(100)
	tx := NewTx(from, to, amount)

	assert.Equal(t, "0x4c9e722b72ffe104c937d62d4e449b0307a474b2ebb808455ffb37d7ad9e069a", tx.Hash.String())
	assert.Equal(t, "cc56f727b11289925941d9100b992b0e60ff1c8e40ef59c06031bd167ed206f404f2a7f38b377df964000000000000004c9e722b72ffe104c937d62d4e449b0307a474b2ebb808455ffb37d7ad9e069a", tx.Hex())
}

func TestLevel(t *testing.T) {
	from := common.HexToAddress("0xcc56f727B11289925941D9100b992b0e60ff1C8e")
	to := common.HexToAddress("0x40eF59C06031BD167Ed206F404f2A7f38b377Df9")
	amount := uint64(100)
	tx := NewTx(from, to, amount)

	// open db
	db, err := leveldb.OpenFile("./tmp", nil)
	require.Nil(t, err)

	// write db
	err = db.Put([]byte("key"), tx.Bytes(), nil)
	require.Nil(t, err)

	// read db
	data, err := db.Get([]byte("key"), nil)
	require.Nil(t, err)

	tx2 := NewTxFromBytes(data)

	assert.Equal(t, tx, tx2)

	defer db.Close()
}
