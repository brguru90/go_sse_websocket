package api_module

import (
	"io"

	"net/url"

	"github.com/gin-gonic/gin"
)

func InitSSE(msgGeneratorFunc func(c *gin.Context, stream chan string)) func(c *gin.Context) {
	return func(c *gin.Context) {
		stream := make(chan string)
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")

		go msgGeneratorFunc(c.Copy(), stream)
		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-stream; ok {
				// c.Writer.Write([]byte(msg)) // may be can be used in streaming binary data like files
				// c.SSEvent("message", msg) // method in official documentation
				c.Writer.Write([]byte("data: " + url.PathEscape(msg) + "\n\n"))
				return true
			} else {
				c.Writer.Flush()
			}

			return false
		})
	}
}
