package main

import (
	"log"

	"github.com/eron97/api_createAccount/config/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
