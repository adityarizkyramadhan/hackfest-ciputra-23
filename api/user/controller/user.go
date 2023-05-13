package controller

import (
	"net/http"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/api/user/usecase"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/middleware"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/customserror"
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
	response.Success(ctx, http.StatusCreated, gin.H{"token": token})
}

func (ctrl *User) AddLocation(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	if id == "" {
		response.Fail(ctx, http.StatusForbidden, customserror.ErrIdNotFound.Error())
		return
	}
	var input model.UserLocationRequest
	if err := ctx.BindJSON(&input); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err := ctrl.usecaseUser.AddLocation(&input, id); err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, nil)
}

func (ctrl *User) IsUserAddLocation(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	if id == "" {
		response.Fail(ctx, http.StatusForbidden, customserror.ErrIdNotFound.Error())
		return
	}
	status, err := ctrl.usecaseUser.IsUserAddLocation(id)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, gin.H{"has_pick_location": status})
}

func (ctrl *User) UserLocation(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	if id == "" {
		response.Fail(ctx, http.StatusForbidden, customserror.ErrIdNotFound.Error())
		return
	}
	location, err := ctrl.usecaseUser.GetUserLocation(id)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, location)
}

func (ctrl *User) Profile(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	if id == "" {
		response.Fail(ctx, http.StatusForbidden, customserror.ErrIdNotFound.Error())
		return
	}
	location, err := ctrl.usecaseUser.GetUserById(id)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, location)
}

func (ctrl *User) Mount(user *gin.RouterGroup) {
	user.POST("login", ctrl.Login)
	user.POST("register", ctrl.Register)
	user.POST("location", middleware.ValidateJWToken(), ctrl.AddLocation)
	user.GET("location", middleware.ValidateJWToken(), ctrl.UserLocation)
	user.GET("location/status", middleware.ValidateJWToken(), ctrl.IsUserAddLocation)
	user.GET("profile", middleware.ValidateJWToken(), ctrl.Profile)
}
