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

func StartUserStream(ctx context.Context, key *pub.Key) (*websocket.Conn, chan interface{}, error) {
	if key == nil || key.ApiKey == "" || key.SecretKey == "" {
		return nil, nil, fmt.Errorf("key is nil or api key, secret key is empty")
	}

	var err error
	listenKey, err := GetListenKey(key)
	if err != nil {
		return nil, nil, err
	}
	go func() { // to refrest listen key in each 58 minutes, because listen key will expire in 60 minutes
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

	urlPath := "/ws/" + listenKey
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
				return // return not break
			case msg := <-rawDataChan:
				if msg == nil {
					continue
				}
				data, err := userDataProcess(key, msg)
				if err != nil {
					log.Printf("StartUserStream userDataProcess err: %v, msg:%v", err, string(msg.Message))
					continue
				}
				if data != nil {
					if str, ok := data.(string); ok {
						if str == "listenKeyExpired" {
							listenKey, _ = GetListenKey(key)
							return // exit reading, and will reconnect
						}
					} else {
						processedDataChan <- data
					}
				}

			}
		}
	}()

	return conn, processedDataChan, nil
}

func userDataProcess(key *pub.Key, data *pub.WsMessage) (interface{}, error) {
	if key == nil || data == nil || data.Message == nil {
		return nil, fmt.Errorf("data is nil or message is nil")
	}

	fmt.Printf("userDataProcess: %+v\n", string(data.Message))
	var d streamHeader
	if err := json.Unmarshal(data.Message, &d); err != nil {
		return nil, err
	}
	switch d.EventType {
	case "listenKeyExpired":
		var m streamHeader
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		fmt.Printf("listenKeyExpired: %+v\n", m)
		return "listenKeyExpired", nil

	case "ACCOUNT_UPDATE":
		var m AccountUpdate
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		m.UserId = key.UserId
		return m, nil
	case "MARGIN_CALL":
		var m MarginCall
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		m.UserId = key.UserId
		return m, nil

	case "ORDER_TRADE_UPDATE":
		var m OrderTradeUpdate
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		m.UserId = key.UserId
		return m, nil

	case "TRADE_LITE":
		var m TradeLite
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		m.UserId = key.UserId
		return m, nil

	case "ACCOUNT_CONFIG_UPDATE":
		var m AccountConfigUpdate
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		m.UserId = key.UserId
		return m, nil

	case "STRATEGY_UPDATE":
		var m StrategyUpdate
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		m.UserId = key.UserId
		return m, nil

	case "GRID_UPDATE":
		var m GridUpdate
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		m.UserId = key.UserId
		return m, nil

	case "CONDITIONAL_ORDER_TRIGGER_REJECT":
		var m ConditionalOrderTriggerReject
		if err := json.Unmarshal(data.Message, &m); err != nil {
			return nil, err
		}
		m.UserId = key.UserId
		return m, nil

	default:
		return nil, fmt.Errorf("unknown event type: %s", d.EventType)
	}
}

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
