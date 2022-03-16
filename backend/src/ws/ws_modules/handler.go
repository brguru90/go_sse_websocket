package ws_modules

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type WsHandlers struct {
	M *melody.Melody
}

var AllHandler map[string]WsHandlers = make(map[string]WsHandlers)

func GetWsHandlers(handler_id string, api_callback func(*gin.Context, WsHandlers), handler_callback func(*melody.Session, []byte, WsHandlers)) func(c *gin.Context) {
	if AllHandler[handler_id].M != nil {
		var WS WsHandlers
		WS.M = melody.New()
		WS.M.HandleMessage(func(s *melody.Session, b []byte) {
			handler_callback(s, b, WS)
		})
		AllHandler[handler_id] = WS
	}
	return func(c *gin.Context) {
		api_callback(c, AllHandler[handler_id])
	}
}
