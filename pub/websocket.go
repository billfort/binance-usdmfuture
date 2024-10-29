package pub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

type WsMessage struct {
	MsgType int
	Message []byte
}
type StreamDataCallback func(msg *WsMessage) error

func WsConnect(ctx context.Context, urlPath string, data chan *WsMessage) (*websocket.Conn, error) {
	url := endpoint_websocket + urlPath
	fmt.Println("WsSubscribe url:", url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		defer conn.Close()

		for {
			if ctx.Err() != nil {
				log.Printf("OpenWebsocketConn break read loop now because of context err: %v", ctx.Err())
				break
			}
			msgType, message, err := conn.ReadMessage()
			if err != nil {
				if !strings.Contains(err.Error(), "normal") { // websocket: close 1000 (normal): Bye
					log.Printf("OpenWebsocketConn ReadMessage websocket err: %v, exit now.", err)
				}
				return
			}
			// fmt.Printf("websocket got: msgType: %v, message: %v\n", msgType, string(message))
			if msgType == websocket.PingMessage {
				conn.WriteMessage(websocket.PongMessage, nil)
				continue
			}

			var errmsg ErrMsg
			err = json.Unmarshal(message, &errmsg)
			if err != nil {
				log.Printf("OpenWebsocketConn json.Unmarshal err %v, msg:%v", err, message)
				continue
			}

			if errmsg.Code < 0 { // error happen {code: -xxxx, message: ... }
				log.Printf("OpenWebsocketConn websocket got err message: %+v", errmsg)
				return
			}

			select {
			case data <- &WsMessage{MsgType: msgType, Message: message}:
			case <-ctx.Done():
				log.Printf("OpenWebsocketConn break read loop now because of context err: %v", ctx.Err())
				return
			}
		}
	}()

	return conn, nil
}
