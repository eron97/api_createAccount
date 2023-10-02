package controller

import (
	"github.com/eron97/api_createAccount/config/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Rota chamada de forma errada")
	c.JSON(err.Code, err)

}
