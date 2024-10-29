package streammarket

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/billfort/binance-usdmfuture/pub"
)

func aggTradeStreamDataProcess(data *pub.WsMessage) error {
	var t AggTrade
	if err := json.Unmarshal(data.Message, &t); err != nil {
		return err
	}
	fmt.Printf("aggTrade: %+v\n", t)

	return nil
}

func multiStreamDataProcess(data *pub.WsMessage) error {
	var d StreamData
	if err := json.Unmarshal(data.Message, &d); err != nil {
		return err
	}
	fmt.Printf("StreamData: %+v\n", d.Stream)
	if d.Stream == "" || d.Data == nil {
		return fmt.Errorf("Got non-stream data: %s\n", string(data.Message))
	}

	b, err := json.Marshal(d.Data)
	if err != nil {
		return err
	}

	eventyType := d.Data["e"].(string)
	switch eventyType {
	case "aggTrade":
		var t AggTrade
		if err := json.Unmarshal(b, &t); err != nil {
			return err
		}
		// fmt.Printf("aggTrade: %+v\n", t)

	case "markPriceUpdate":
		var m MarkPriceUpdate
		if err := json.Unmarshal(b, &m); err != nil {
			return err
		}
		fmt.Printf("markPriceUpdate: %+v\n", m)
	case "kline":
		var k Kline
		if err := json.Unmarshal(b, &k); err != nil {
			return err
		}
		// fmt.Printf("kline: %+v\n", k)
	}

	return nil
}

// go test -v -run TestSubscribeSingleStream
func TestSubscribeSingleStream(t *testing.T) {
	streams := []string{"btcusdt@aggTrade"}
	ctx, cancel := context.WithCancel(context.Background())
	_, done, err := StartSubscribe(ctx, streams, aggTradeStreamDataProcess)
	if err != nil {
		t.Error(err)
	}
	select {
	case <-done:
		fmt.Printf("done")
	case <-time.After(10 * time.Second):
		t.Log("Run 10 seconds, quit now.")
	}
	cancel()
}

// go test -v -run TestSubscribeMultiStream
func TestSubscribeMultiStream(t *testing.T) {
	streams := []string{"btcusdt@aggTrade", "ethusdt@kline_1h"}
	ctx, cancel := context.WithCancel(context.Background())
	_, done, err := StartSubscribe(ctx, streams, multiStreamDataProcess)
	if err != nil {
		t.Error(err)
	}
	select {
	case <-done:
		fmt.Printf("done")
	case <-time.After(10 * time.Second):
		t.Log("Run 10 seconds, quit now.")
	}
	cancel()
}

// go test -v -run TestSubUnsub
func TestSubUnsub(t *testing.T) {
	streams := []string{"btcusdt@aggTrade", "ethusdt@aggTrade"}
	ctx, cancel := context.WithCancel(context.Background())
	conn, done, err := StartSubscribe(ctx, streams, multiStreamDataProcess)
	if err != nil {
		t.Error(err)
		cancel()
		return
	}
	go func() { // sub
		<-time.After(3 * time.Second)
		streams = []string{"btcusdt@kline_1h"}
		if err := SubUnSub(conn, streams, "SUBSCRIBE"); err != nil {
			t.Error(err)
		} else {
			fmt.Println("Subscribed")
		}
	}()
	go func() { // unsub
		<-time.After(8 * time.Second)
		streams = []string{"btcusdt@aggTrade"}
		if err := SubUnSub(conn, streams, "UNSUBSCRIBE"); err != nil {
			t.Error(err)
		} else {
			fmt.Println("Unsubscribed")
		}
	}()

	select {
	case <-done:
		fmt.Printf("done")
	case <-time.After(10 * time.Second):
		t.Log("Run 10 seconds, quit now.")
	}
	cancel()
}
