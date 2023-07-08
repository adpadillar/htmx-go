package routes

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.File("src/static/index.html")
}
