package views

import (
	"fmt"
	"learn_go/src/ws/ws_modules"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func TestWS(c *gin.Context, w ws_modules.WsHandlers) {
	fmt.Println("TestWS")
	w.M.HandleRequest(c.Writer, c.Request)
}

func TestWSH(s *melody.Session, msg []byte, w ws_modules.WsHandlers) {
	switch string(msg) {
	case "hi":
		w.M.Broadcast([]byte("bi"))
	case "GM":
		w.M.Broadcast([]byte("GN"))
	default:
		w.M.Broadcast(msg)
	}

}
