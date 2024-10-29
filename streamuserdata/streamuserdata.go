package streamuserdata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/billfort/binance-usdmfuture/pub"
	"github.com/gorilla/websocket"
)

func GetListenKey(key *pub.Key) (string, error) {
	resBody, errMsg, err := pub.PostWithSign(key, "/fapi/v1/listenKey", nil)
	if err != nil {
		return "", err
	}
	if errMsg.Code < 0 {
		return "", fmt.Errorf("%+v", errMsg)
	}

	var lk listenKey
	if err := json.Unmarshal(resBody, &lk); err != nil {
		return "", err
	}

	return lk.ListenKey, nil
}

// PutListenKey updates the listen key. keep alive in 60 minutes.
func PutListenKey(key *pub.Key) (string, error) {
	resBody, err := pub.PutWithSign(key, "/fapi/v1/listenKey", nil)
	if err != nil {
		return "", err
	}

	var lk listenKey
	if err := json.Unmarshal(resBody, &lk); err != nil {
		return "", err
	}

	return lk.ListenKey, nil
}

// DeleteListenKey deletes the listen key.
func DeleteListenKey(key *pub.Key) error {
	_, err := pub.DeleteWithSign(key, "/fapi/v1/listenKey", nil)
	if err != nil {
		return err
	}

	return nil
}

func StartUserStream(ctx context.Context, key *pub.Key, cb pub.StreamDataCallback) (*websocket.Conn, chan bool, error) {
	listenKey, err := GetListenKey(key)
	if err != nil {
		return nil, nil, err
	}

	urlPath := "/ws/" + listenKey
	fmt.Println("urlPath:", urlPath)
	dataStream := make(chan *pub.WsMessage, 100)
	conn, err := pub.WsConnect(ctx, urlPath, dataStream)
	if err != nil {
		return nil, nil, err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("StartUserStream break now because of context err: %v", ctx.Err())
				return
			case <-time.After(58 * time.Minute): // keey alive each 60 minutes
				_, err := PutListenKey(key)
				if err != nil {
					log.Printf("StartUserStream PutListenKey err: %v", err)
					return
				}
			}
		}
	}()

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
						log.Printf("StartUserStream callback err: %v", err)
					}
				}
			}
		}
	}()

	return conn, done, nil
}
