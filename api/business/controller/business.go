package controller

import (
	"net/http"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/api/business/usecase"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/middleware"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Business struct {
	usecaseBusines usecase.Business
}

func New(usecaseBusines usecase.Business) *Business {
	return &Business{usecaseBusines}
}

func (ctrl *Business) CreateBusiness(ctx *gin.Context) {
	input := new(model.BusinessInput)
	if err := ctx.ShouldBindWith(input, binding.FormMultipart); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err := ctrl.usecaseBusines.CreateBusiness(input); err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, nil)
}

func (ctrl *Business) CreateTestimony(ctx *gin.Context) {
	var input model.TestimonyInput
	idUser := ctx.MustGet("id").(string)
	if err := ctx.Bind(&input); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err := ctrl.usecaseBusines.CreateTestimony(&input, idUser); err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, nil)
}

func (ctrl *Business) GetBusiness(ctx *gin.Context) {
	tipe := ctx.Query("type")
	offered := ctx.Query("offered")
	data, err := ctrl.usecaseBusines.GetBussiness(tipe, offered)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, data)
}

func (ctrl *Business) GetBusinessById(ctx *gin.Context) {
	data, err := ctrl.usecaseBusines.GetByIdBusiness(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, data)
}

func (ctrl *Business) Mount(business *gin.RouterGroup) {
	business.GET(":id", middleware.ValidateJWToken(), ctrl.GetBusinessById)
	business.POST("seller", ctrl.CreateBusiness)
	business.POST("comment", middleware.ValidateJWToken(), ctrl.CreateTestimony)
	business.GET("query", middleware.ValidateJWToken(), ctrl.GetBusiness)
}
