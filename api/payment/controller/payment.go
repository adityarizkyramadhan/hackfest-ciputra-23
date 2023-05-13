package controller

import (
	"net/http"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/api/user/usecase"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/middleware"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/payment"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/response"
	"github.com/gin-gonic/gin"
)

type Payment struct {
	usecaseUser *usecase.User
}

func New(usecaseUser *usecase.User) *Payment {
	return &Payment{usecaseUser}
}

func (ctrl *Payment) CreatePayment(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	var request payment.PaymentRequest
	if err := ctx.BindJSON(&request); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	user, err := ctrl.usecaseUser.GetUserById(id)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	invoiceXendit, err := payment.CreatePayment(&request, user)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, invoiceXendit)
}

func (ctrl *Payment) Success(ctx *gin.Context) {
	response.Success(ctx, http.StatusOK, gin.H{"status_payment": "success"})
}

func (ctrl *Payment) Fail(ctx *gin.Context) {
	response.Fail(ctx, http.StatusOK, "payment fail")
}

func (ctrl *Payment) Mount(payment *gin.RouterGroup) {
	payment.POST("pay", middleware.ValidateJWToken(), ctrl.CreatePayment)
	payment.GET("success", ctrl.Success)
	payment.GET("failure", ctrl.Fail)
}
