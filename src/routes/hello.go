package routes

import (
	"github.com/gin-gonic/gin"
)

const ContentTypeHTML = "text/html; charset=utf-8"

func HelloHandler(c *gin.Context) {
	const html = `<h1>Hello, world!!!</h1>`

	c.Data(200, ContentTypeHTML, []byte(html))
}
