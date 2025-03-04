package streammarket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/billfort/binance-usdmfuture/pub"
	"github.com/gorilla/websocket"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func StartSubscribe(ctx context.Context, streams []string) (*websocket.Conn, chan interface{}, error) {
	var urlPath string
	if len(streams) == 1 {
		urlPath = "/ws/" + streams[0]
	} else {
		urlPath = "/stream?streams=" + strings.Join(streams, "/")
	}
	fmt.Println("stream market urlPath:", urlPath)

	conn, rawDataChan, err := pub.WsConnect(ctx, urlPath)
	if err != nil {
		log.Printf("StartSubscribe WsConnect err: %v", err)
		return nil, nil, err
	}

	processedDataChan := make(chan interface{}, pub.WsChanLen)
	go func() {
		defer close(processedDataChan)

		for { // read message loop
			select {
			case <-ctx.Done():
				return // return, no more read.
			case msg := <-rawDataChan:
				if msg.MsgType == websocket.CloseMessage { // connection closed
					log.Printf("StartSubscribe close message: %+v", msg)
					return // return, no more read.
				}

				d, err := streamDataProcess(msg)
				if err != nil {
					log.Printf("StartSubscribe streamDataProcess err: %v", err)
				}
				if d != nil {
					processedDataChan <- d
				}
			}
		}
	}()

	return conn, processedDataChan, nil
}

func streamDataProcess(m *pub.WsMessage) (interface{}, error) {
	var d StreamData
	if err := json.Unmarshal(m.Message, &d); err != nil {
		return nil, err
	}
	if d.Stream == "" && d.Data == nil { // single stream data
		if err := json.Unmarshal(m.Message, &d.Data); err != nil {
			return nil, err
		}
	}

	if d.Stream == "" && d.Data == nil {
		return nil, fmt.Errorf("got non-stream data: %s", string(m.Message))
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
		return t, nil

	case "markPriceUpdate":
		var t MarkPriceUpdate
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "kline":
		var t Kline
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "miniTicker":
		var t MiniTicker
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "ticker":
		var t Ticker
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "bookTicker":
		var t BookTicker
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "forceOrder":
		var t ForceOrder
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "depthUpdate":
		var t DepthUpdate
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "compositeIndex":
		var t CompositeIndex
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "contractInfo":
		var t ContractInfo
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	case "assetIndexUpdate":
		var t AssetIndexUpdate
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, err
		}
		return t, nil

	default:
		return nil, fmt.Errorf("unknown event type: %s, data: %s", eventyType, string(m.Message))
	}
}

// method: SUBSCRIBE, UNSUBSCRIBE
func SubUnSub(conn *websocket.Conn, streams []string, method string) error {
	var sub SubUnsub
	sub.Method = method
	sub.Params = streams
	sub.ID = time.Now().UnixMilli()
	b, err := json.Marshal(sub)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}

	return nil
}
