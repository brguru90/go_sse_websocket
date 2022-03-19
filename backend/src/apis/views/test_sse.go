package views

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TestSSE(c *gin.Context, stream chan string) {
	for i := 0; i < 10; i++ {
		stream <- fmt.Sprintf("chunk %d", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(stream)
}
