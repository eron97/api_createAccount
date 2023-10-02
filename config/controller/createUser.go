package controller

import (
	"fmt"
	"log"

	"github.com/eron97/api_createAccount/config/controller/model/request"
	"github.com/eron97/api_createAccount/config/validation"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal objetct, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
