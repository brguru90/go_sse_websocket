package views

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"learn_go/src/ws/ws_modules"
)


func TestWS(c *gin.Context,w ws_modules.WsHandlers) {	
	w.M.HandleRequest(c.Writer, c.Request)
}

func  TestWSH(s *melody.Session, msg []byte,w ws_modules.WsHandlers) {
	w.M.Broadcast(msg)
}
