package streammarket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/billfort/binance-usdmfuture/pub"
	"github.com/gorilla/websocket"
)

func StartSubscribe(ctx context.Context, streams []string, cb pub.StreamDataCallback) (*websocket.Conn, chan bool, error) {
	var urlPath string
	if len(streams) == 1 {
		urlPath = "/ws/" + streams[0]
	} else {
		urlPath = "/stream?streams=" + strings.Join(streams, "/")
	}
	fmt.Println("urlPath:", urlPath)
	dataStream := make(chan *pub.WsMessage, 100)
	conn, err := pub.WsConnect(ctx, urlPath, dataStream)
	if err != nil {
		return nil, nil, err
	}

	done := make(chan bool)
	go func() {
		defer conn.Close()
		defer close(done)

		for {
			select {
			case <-ctx.Done():
				return // return not break
			case msg := <-dataStream:
				if cb != nil {
					err = cb(msg)
					if err != nil {
						log.Printf("StartSubscribe callback err: %v", err)
					}
				}
			}
		}
	}()

	return conn, done, nil
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
