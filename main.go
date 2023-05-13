package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllerUser "github.com/adityarizkyramadhan/hackfest-ciputra-23/api/user/controller"
	repositoryUser "github.com/adityarizkyramadhan/hackfest-ciputra-23/api/user/repository"
	usecaseUser "github.com/adityarizkyramadhan/hackfest-ciputra-23/api/user/usecase"

	controllerPayment "github.com/adityarizkyramadhan/hackfest-ciputra-23/api/payment/controller"

	controllerBusiness "github.com/adityarizkyramadhan/hackfest-ciputra-23/api/business/controller"
	repositoryBusiness "github.com/adityarizkyramadhan/hackfest-ciputra-23/api/business/repository"
	usecaseBusiness "github.com/adityarizkyramadhan/hackfest-ciputra-23/api/business/usecase"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/config/database"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/middleware"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/model"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
	db := database.Init()

	if db == nil {
		log.Fatal("init connection db failed")
	}
	err = database.Migrate(
		&model.User{},
		&model.UserLocation{},
		&model.Business{},
		&model.Testimony{},
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	router := gin.New()
	router.Use(middleware.CORS())
	router.Use(middleware.Timeout(30))
	router.GET("health", func(ctx *gin.Context) {
		response.Success(ctx, http.StatusOK, "api healty 100%")
	})

	api := router.Group("api")
	v1 := api.Group("v1")

	repoUser := repositoryUser.New(db)
	ucUser := usecaseUser.New(repoUser)
	ctrlUser := controllerUser.New(ucUser)
	userGroup := v1.Group("user")
	ctrlUser.Mount(userGroup)

	repoBusiness := repositoryBusiness.New(db)
	ucBusiness := usecaseBusiness.New(repoBusiness)
	ctrlBusiness := controllerBusiness.New(*ucBusiness)
	businessGroup := v1.Group("business")
	ctrlBusiness.Mount(businessGroup)

	ctrlPayment := controllerPayment.New(ucUser)
	paymentGroup := v1.Group("payment")
	ctrlPayment.Mount(paymentGroup)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
