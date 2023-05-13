package controller

import (
	"net/http"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/api/user/usecase"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	usecaseUser *usecase.User
}

func New(usecaseUser *usecase.User) *User {
	return &User{usecaseUser}
}

func (ctrl *User) Login(ctx *gin.Context) {
	var input model.UserLogin
	if err := ctx.BindJSON(&input); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	token, err := ctrl.usecaseUser.Login(&input)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, gin.H{"token": token})
}

func (ctrl *User) Register(ctx *gin.Context) {
	var input model.UserRegister
	if err := ctx.BindJSON(&input); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	token, err := ctrl.usecaseUser.Register(&input)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, gin.H{"token": token})
}

func (ctrl *User) Mount(user *gin.RouterGroup) {
	user.POST("login", ctrl.Login)
	user.POST("register", ctrl.Register)
}
