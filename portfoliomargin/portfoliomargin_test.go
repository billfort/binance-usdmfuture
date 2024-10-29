package portfoliomargin

import (
	"fmt"
	"testing"

	"github.com/billfort/binance-usdmfuture/pub"
	"github.com/stretchr/testify/require"
)

// go test -v -run TestGetPmAccountInfo
func TestGetPmAccountInfo(t *testing.T) {
	res, err := GetPmAccountInfo(pub.TestKey, "BTC")
	require.Nil(t, err)
	require.NotNil(t, res)
	fmt.Printf("GetPmAccountInfo: %+v\n", res)
}
