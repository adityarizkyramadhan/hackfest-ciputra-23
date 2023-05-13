package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/config/database"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/middleware"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
	router := gin.New()

	router.Use(middleware.CORS())

	db := database.Init()

	if db == nil {
		log.Fatal("init connection db failed")
	}

	router.Use(middleware.Timeout(30))
	router.GET("health", func(ctx *gin.Context) {
		response.Success(ctx, http.StatusOK, "api healty 100%")
	})
	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
