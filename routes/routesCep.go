package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/guilherm5/cep/controllers"
)

func CepGet(c *gin.Engine) {
	c.POST("/cep", controllers.ObtemCep)
}
