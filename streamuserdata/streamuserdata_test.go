package streamuserdata

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/billfort/binance-usdmfuture/pub"
	"github.com/stretchr/testify/require"
)

// go test -v -run TestGetListenKey
func TestGetListenKey(t *testing.T) {
	listenKey, err := GetListenKey(pub.TestKey)
	require.NoError(t, err)
	require.NotEmpty(t, listenKey)
}

// go test -v -run TestStartUserStream
func TestStartUserStream(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	conn, data, err := StartUserStream(ctx, pub.TestKey)
	require.NoError(t, err)
	require.NotNil(t, conn)

Loop:
	for {
		select {
		case d := <-data:
			fmt.Printf("data: %+v\n", d)
		case <-time.After(10 * time.Minute):
			t.Log("Run 10 minutes, quit now.")
			break Loop
		}
	}
	cancel()
}
