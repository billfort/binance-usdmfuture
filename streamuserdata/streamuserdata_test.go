package streamuserdata

import (
	"context"
	"encoding/json"
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

func userDataProcess(data *pub.WsMessage) error {
	var d streamHeader
	if err := json.Unmarshal(data.Message, &d); err != nil {
		return err
	}
	switch d.EventType {
	case "listenKeyExpired":
		var t streamHeader
		if err := json.Unmarshal(data.Message, &t); err != nil {
			return err
		}
		fmt.Printf("listenKeyExpired: %+v\n", t)

	case "ACCOUNT_UPDATE":
		var m AccountUpdate
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return err
		}
		fmt.Printf("AccountUpdate: %+v\n", m)
	case "ORDER_TRADE_UPDATE":
		var k OrderTradeUpdate
		if err := json.Unmarshal(data.Message, &k); err != nil {
			return err
		}
		fmt.Printf("ORDER_TRADE_UPDATE: %+v\n", k)
	case "TRADE_LITE":
		var k TradeLite
		if err := json.Unmarshal(data.Message, &k); err != nil {
			return err
		}
		fmt.Printf("TRADE_LITE: %+v\n", k)
	}

	return nil
}

// go test -v -run TestStartUserStream
func TestStartUserStream(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	conn, done, err := StartUserStream(ctx, pub.TestKey, userDataProcess)
	require.NoError(t, err)
	require.NotNil(t, conn)

	select {
	case <-done:
		fmt.Printf("done")
	case <-time.After(2 * time.Minute):
		t.Log("Run 10 seconds, quit now.")
	}
	cancel()
}
