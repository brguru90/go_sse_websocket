package ws

import (
	"learn_go/src/ws/views"

	"github.com/gin-gonic/gin"
)

func InitWS(router *gin.RouterGroup) {
	router.GET("chat/", views.TestWSH)
}
