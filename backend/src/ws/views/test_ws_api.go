package views

import (
	"fmt"
	"learn_go/src/ws/ws_modules"

	"gopkg.in/olahol/melody.v1"
)

func TestWSH(w *ws_modules.WsHandlers) {

	w.OnConnect = func(s *melody.Session) {
		fmt.Println("on Connect")
		w.M.Broadcast([]byte("connection success : initial message from server"))
	}

	w.OnMessage = func(s *melody.Session, msg []byte) {
		fmt.Println("on Receiving message")
		switch string(msg) {
		case "hi":
			w.M.Broadcast([]byte("bi"))
		case "GM":
			w.M.Broadcast([]byte("GN"))
		default:
			w.M.Broadcast(msg)
		}
	}

	w.OnDisconnect=func(s *melody.Session) {
		fmt.Println("on Disconnect")
	}

}
