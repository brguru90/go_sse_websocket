package main

import (
	"fmt"
	"os"

	"learn_go/src/apis"
	"learn_go/src/ws"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var all_router *gin.Engine

var SERVER_PORT string = "8000"

func main() {
	// all_router = gin.New()
	all_router = gin.Default()
	all_router.Use(cors.Default())
	all_router.Use(static.Serve("/", static.LocalFile("./src/static", true)))
	{
		api_router := all_router.Group("/api")
		apis.InitApis(api_router)
	}

	{
		ws_router := all_router.Group("/ws")
		ws.InitWS(ws_router)
	}

	if os.Getenv("SERVER_PORT") != "" {
		SERVER_PORT = os.Getenv("SERVER_PORT")
	}

	bind_to_host := fmt.Sprintf(":%s", SERVER_PORT) //formatted host string
	all_router.Run(bind_to_host)
}

