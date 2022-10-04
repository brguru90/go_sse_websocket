package ws

import (
	"learn_go/src/ws/views"
	"learn_go/src/ws/ws_modules"

	"github.com/gin-gonic/gin"
)

func InitWS(router *gin.RouterGroup) {
	router.GET(ws_modules.GetWsHandlers("chat/",views.TestWSH))
}
