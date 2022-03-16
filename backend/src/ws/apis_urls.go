package ws

import (
	"learn_go/src/ws/views"
	"learn_go/src/ws/ws_modules"

	"github.com/gin-gonic/gin"
)



func InitWS(router *gin.RouterGroup) {
	router.GET("chat/", ws_modules.GetWsHandlers("test_ws",views.TestWS,views.TestWSH))
}
