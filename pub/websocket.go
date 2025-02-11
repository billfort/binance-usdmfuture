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

func WsConnect(ctx context.Context, urlPath string) (*websocket.Conn, chan *WsMessage, error) {
	url := futureWssUrl + urlPath
	fmt.Println("WsConnect url:", url)

	if ctx.Err() != nil {
		log.Printf("WsConnect context err: %v", ctx.Err())
		return nil, nil, ctx.Err()
	}
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("WsConnect websocket dial %s err: %v", url, err)
		return nil, nil, err
	}

	rawDataChan := make(chan *WsMessage, WsChanLen)

	go func() { // read message loop
		defer func() {
			rawDataChan <- &WsMessage{MsgType: websocket.CloseMessage, Message: nil}
			conn.Close()
			close(rawDataChan)
			fmt.Printf("WsConnect read message loop exit now.\n")
		}()

		for {
			if ctx.Err() != nil {
				log.Printf("OpenWebsocketConn break read loop now because of context err 1: %v", ctx.Err())
				return
			}

			msgType, message, err := conn.ReadMessage()
			if err != nil {
				if !strings.Contains(err.Error(), "normal") { // websocket: close 1000 (normal): Bye
					log.Printf("OpenWebsocketConn ReadMessage websocket err: %v, exit now.", err)
				}
				return
			}

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
			case rawDataChan <- &WsMessage{MsgType: msgType, Message: message}:
			case <-ctx.Done():
				log.Printf("OpenWebsocketConn break read loop now because of context err 2: %v", ctx.Err())
				return
			}
		}
	}()

	return conn, rawDataChan, nil
}
