package views

import (
	"fmt"
	"learn_go/src/ws/ws_modules"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)


func TestWS(c *gin.Context,w ws_modules.WsHandlers) {	
	fmt.Println("TestWS")
	w.M.HandleRequest(c.Writer, c.Request)
}

func  TestWSH(s *melody.Session, msg []byte,w ws_modules.WsHandlers) {
	fmt.Println("TestWSH",msg)
	w.M.Broadcast(msg)
}
