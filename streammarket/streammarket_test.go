package streammarket

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/billfort/binance-usdmfuture/pub"
)

// go test -v -run TestSubscribeSingleStream
func TestSubscribeSingleStream(t *testing.T) {
	streams := []string{"btcusdt@aggTrade"}
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-time.After(8 * time.Second)
		t.Log("Run 8 seconds, quit now.")
		cancel()
	}()

	for { // loop to test if connection closed
		conn, ch, err := StartSubscribe(ctx, streams)
		if err != nil {
			t.Error(err)
			break // you can sleep some time and continue here.
		}

		go func() { // close connection after 3 seconds
			<-time.After(3 * time.Second)
			t.Log("Run 3 seconds, close conn now.")
			conn.Close()
		}()

		for { // read message loop
			data := <-ch
			fmt.Printf("data: %+v\n", data)
			if data == nil {
				break
			}
		}
	}
}

// go test -v -run TestSubscribeMultiStream
func TestSubscribeMultiStream(t *testing.T) {
	streams := []string{"btcusdt@aggTrade", "ethusdt@kline_1h"}
	ctx, cancel := context.WithCancel(context.Background())
	_, ch, err := StartSubscribe(ctx, streams)
	if err != nil {
		t.Error(err)
	}

	go func() {
		<-time.After(8 * time.Second)
		t.Log("Run 8 seconds, quit now.")
		cancel()
	}()

	for { // read message loop
		data := <-ch
		fmt.Printf("data: %+v\n", data)
		if data == nil {
			break
		}
	}
}

// go test -v -run TestSubUnsub
func TestSubUnsub(t *testing.T) {
	streams := []string{"btcusdt@aggTrade", "ethusdt@aggTrade"}
	ctx, cancel := context.WithCancel(context.Background())
	conn, data, err := StartSubscribe(ctx, streams)
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
	case d := <-data:
		fmt.Printf("data: %+v\n", d)
	case <-time.After(10 * time.Second):
		t.Log("Run 10 seconds, quit now.")
	}
	cancel()
}

func aggTradeStreamDataProcess(data *pub.WsMessage) (interface{}, error) {
	var t AggTrade
	if err := json.Unmarshal(data.Message, &t); err != nil {
		return nil, err
	}
	fmt.Printf("aggTrade: %+v\n", t)

	return nil, nil
}

func multiStreamDataProcess(data *pub.WsMessage) (interface{}, error) {
	var d StreamData
	if err := json.Unmarshal(data.Message, &d); err != nil {
		return nil, err
	}

	if d.Stream == "" || d.Data == nil {
		return nil, fmt.Errorf("Got non-stream data: %s\n", string(data.Message))
	}

	b, err := json.Marshal(d.Data)
	if err != nil {
		return nil, err
	}

	eventyType := d.Data["e"].(string)
	switch eventyType {
	case "aggTrade":
		var t AggTrade
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		// fmt.Printf("aggTrade: %+v\n", t)

	case "markPriceUpdate":
		var m MarkPriceUpdate
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		fmt.Printf("markPriceUpdate: %+v\n", m)
	case "kline":
		var k Kline
		if err := json.Unmarshal(b, &k); err != nil {
			return nil, err
		}
		// fmt.Printf("kline: %+v\n", k)
	}

	return nil, err
}
