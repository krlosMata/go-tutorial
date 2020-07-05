package basic

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	account := NewAccountRandom()

	bytesTx, err := hex.DecodeString("cc56f727b11289925941d9100b992b0e60ff1c8e40ef59c06031bd167ed206f404f2a7f38b377df9640000000000000000128435dfc1389f0236c7e279a6b21c68d999389f1e8857fe0f2d9fe33d5d8600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c1dff5e000000004c01000000000000")
	require.Nil(t, err)
	tx := NewTxFromBytes(bytesTx)

	account.SignTx(&tx)
	require.Nil(t, err)
}
