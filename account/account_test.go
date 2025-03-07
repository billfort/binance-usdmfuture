package account

import (
	"fmt"
	"testing"

	"github.com/billfort/binance-usdmfuture/pub"
	"github.com/stretchr/testify/require"
)

// go test -v -run TestAccountBalance
func TestAccountBalance(t *testing.T) {
	res, err := AccountBalance(pub.TestKey, "v2")
	require.Nil(t, err)
	require.NotNil(t, res)
	fmt.Printf("AccountBalance: %+v\n", res)
}

// go test -v -run TestGetInternalTransferHist
func TestGetInternalTransferHist(t *testing.T) {
	res, err := GetInternalTransferHist(pub.TestKey, "", 0)
	require.Nil(t, err)
	require.NotNil(t, res)
	fmt.Printf("GetInternalTransferHist: %+v\n", res)
}
