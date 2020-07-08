package basic

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewService(t *testing.T) {
	basePath, err := os.Getwd()
	require.Nil(t, err)

	pathDb := path.Join(basePath, "tmp")

	s, err := NewService(pathDb)

	assert.NoError(t, err)
	assert.NotEmpty(t, s)

	removeDb(pathDb)
}

func TestUpdateBalances(t *testing.T) {
	basePath, err := os.Getwd()
	require.Nil(t, err)

	pathDb := path.Join(basePath, "tmp")

	s, err := NewService(pathDb)

	amountToAdd := uint64(10)

	address := common.HexToAddress("0x123456789A123456789A123456789A123456789A")

	balanceOld, err := s.GetBalance(address)
	require.Nil(t, err)
	fmt.Println(balanceOld)

	s.UpdateBalances(address, amountToAdd)

	balanceNew, err := s.GetBalance(address)
	require.Nil(t, err)

	assert.Equal(t, amountToAdd, balanceNew)

	removeDb(pathDb)
}

func TestAddTx(t *testing.T) {
	// create account, tx and sign it
	account := NewAccountRandom()
	from := account.EthAddress
	to := common.HexToAddress("0x40eF59C06031BD167Ed206F404f2A7f38b377Df9")
	amount := uint64(100)
	tx := NewTx(from, to, amount)
	account.SignTx(&tx)

	// create new service
	basePath, err := os.Getwd()
	require.Nil(t, err)

	pathDb := path.Join(basePath, "tmp")

	s, err := NewService(pathDb)

	// fill sender with enough balance to perform transaction
	initialBalance := uint64(200)
	s.UpdateBalances(account.EthAddress, initialBalance)

	// add tx to service
	s.AddTx(tx)

	// Get transaction from Db
	txs, err := s.GetTxs(account.EthAddress)

	assert.Equal(t, tx, txs[0])

	removeDb(pathDb)
}

func removeDb(pathDb string) {
	err := os.RemoveAll(pathDb)
	if err != nil {
		log.Fatal(err)
	}
}
