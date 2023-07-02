package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleycastilho95/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/wesleycastilho95/meu-primeiro-crud-go/src/configuration/rest_err/validation"
	"github.com/wesleycastilho95/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/wesleycastilho95/meu-primeiro-crud-go/src/model"
	"github.com/wesleycastilho95/meu-primeiro-crud-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error tryng to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))
}
