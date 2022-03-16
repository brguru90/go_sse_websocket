package views

import (
	"fmt"
	"learn_go/src/ws/ws_modules"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func TestWS(c *gin.Context, w ws_modules.WsHandlers) {
	fmt.Println("on Connect")
	w.M.HandleRequest(c.Writer, c.Request)
	fmt.Println("on Disconnect")
}

func TestWSH(s *melody.Session, msg []byte, w ws_modules.WsHandlers) {
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
