package routes

import (
	"github.com/gin-gonic/gin"
)

func NewPageHandler(c *gin.Context) {
	c.File("src/static/new_page.html")
}
