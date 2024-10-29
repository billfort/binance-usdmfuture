package convert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -run TestListPairs
func TestListPairs(t *testing.T) {
	res, err := ListPairs("BTC", "USDT")
	require.Nil(t, err)
	require.NotNil(t, res)
	fmt.Printf("ListPairs: %+v\n", res)
}
