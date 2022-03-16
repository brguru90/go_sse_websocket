package apis

import (
	"learn_go/src/apis/views"
	"learn_go/src/apis/api_module"
	"github.com/gin-gonic/gin"
)

// only the functions whose initial letter is upper case only those can be exportable from package
func InitApis(router *gin.RouterGroup) {
	router.Use(ApiSpecificMiddleware())
	router.GET("test/:id", test_api)
	router.GET("hello/", views.Hello_api)
	router.GET("test_sse/", api_module.InitSSE(views.TestSSE))
}


