package views

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func TestWSH(c *gin.Context) {
	M := melody.New()
	M.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println("on Receiving message")
		switch string(msg) {
		case "hi":
			M.Broadcast([]byte("bi"))
		case "GM":
			M.Broadcast([]byte("GN"))
		default:
			M.Broadcast(msg)
		}
	})
	M.HandleConnect(func(s *melody.Session) {
		fmt.Println("on Connect")
		M.Broadcast([]byte("connection success : initial message from server"))

	})
	M.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("on Disconnect")
	})

	M.HandleRequest(c.Writer, c.Request)
}
