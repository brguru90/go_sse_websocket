package ws_modules

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

// when every time creating new melody/websocker object,  requires a new unique variable
// & its difficult to group it
// so using bellow function we can easily map handler to the api in apis_urls

type WsHandlers struct {
	M *melody.Melody
}

var AllHandler map[string]WsHandlers = make(map[string]WsHandlers)

func GetWsHandlers(url string, api_callback func(*gin.Context, WsHandlers), handler_callback func(*melody.Session, []byte, WsHandlers)) (string,func(c *gin.Context)) {
	if AllHandler[url].M == nil {
		var WS WsHandlers
		WS.M = melody.New()
		WS.M.HandleMessage(func(s *melody.Session, msg []byte) {
			handler_callback(s, msg, WS)
		})
		AllHandler[url] = WS
	}
	return url,func(c *gin.Context) {
		api_callback(c, AllHandler[url])
	}
}
