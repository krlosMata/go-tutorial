package basic

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path"
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

	assert.Equal(t, byte(0), tx.Hash[0])
}

func TestParserBytes(t *testing.T) {
	from := common.HexToAddress("0xcc56f727B11289925941D9100b992b0e60ff1C8e")
	to := common.HexToAddress("0x40eF59C06031BD167Ed206F404f2A7f38b377Df9")
	amount := uint64(100)
	tx := NewTx(from, to, amount)

	txBytes := tx.Bytes()
	tx2 := NewTxFromBytes(txBytes)
	assert.Equal(t, tx, tx2)
}

func TestLevel(t *testing.T) {
	from := common.HexToAddress("0xcc56f727B11289925941D9100b992b0e60ff1C8e")
	to := common.HexToAddress("0x40eF59C06031BD167Ed206F404f2A7f38b377Df9")
	amount := uint64(100)
	tx := NewTx(from, to, amount)

	// open db
	basePath, err := os.Getwd()
	require.Nil(t, err)
	pathDb := path.Join(basePath, "tmp")

	db, err := leveldb.OpenFile(pathDb, nil)
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

	removeDb(pathDb)
}

func TestHex(t *testing.T) {
	bytesTx, err := hex.DecodeString("cc56f727b11289925941d9100b992b0e60ff1c8e40ef59c06031bd167ed206f404f2a7f38b377df9640000000000000000128435dfc1389f0236c7e279a6b21c68d999389f1e8857fe0f2d9fe33d5d8600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c1dff5e000000004c01000000000000")
	require.Nil(t, err)
	tx := NewTxFromBytes(bytesTx)
	assert.Equal(t, "cc56f727b11289925941d9100b992b0e60ff1c8e40ef59c06031bd167ed206f404f2a7f38b377df9640000000000000000128435dfc1389f0236c7e279a6b21c68d999389f1e8857fe0f2d9fe33d5d8600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c1dff5e000000004c01000000000000", tx.Hex())
}

func TestJson(t *testing.T) {
	from := common.HexToAddress("0xcc56f727B11289925941D9100b992b0e60ff1C8e")
	to := common.HexToAddress("0x40eF59C06031BD167Ed206F404f2A7f38b377Df9")
	amount := uint64(100)
	tx := NewTx(from, to, amount)

	// convert to JSON
	res, err := json.Marshal(tx)
	require.Nil(t, err)

	// load from JSON
	var tx2 Tx
	json.Unmarshal(res, &tx2)

	assert.Equal(t, tx, tx2)
}

func TestVerifyTx(t *testing.T) {
	account := NewAccountRandom()

	from := account.EthAddress
	to := common.HexToAddress("0x40eF59C06031BD167Ed206F404f2A7f38b377Df9")
	amount := uint64(100)
	tx := NewTx(from, to, amount)

	account.SignTx(&tx)

	res := tx.VerifyTx()

	assert.Equal(t, true, res)
}
