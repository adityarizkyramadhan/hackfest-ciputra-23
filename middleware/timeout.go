package middleware

import (
	"net/http"
	"time"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/customserror"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/response"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Timeout(second int) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Duration(second)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			response.Fail(c, http.StatusRequestTimeout, customserror.ErrTimeout.Error())
		}),
	)
}
