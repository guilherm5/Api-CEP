package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/cep/database"
	"github.com/guilherm5/cep/routes"
)

func main() {
	database.Init()
	router := gin.Default()

	routes.CepGet(router)

	router.Run(":2000")
}
