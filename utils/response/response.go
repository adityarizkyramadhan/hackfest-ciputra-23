package response

import "github.com/gin-gonic/gin"

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Success(ctx *gin.Context, status int, data any) {
	meta := Meta{
		Success: true,
		Message: "success",
	}
	response := Response{
		Meta: meta,
		Data: data,
	}
	ctx.JSON(status, response)
}

func Fail(ctx *gin.Context, status int, errorMessage string) {
	meta := Meta{
		Success: false,
		Message: errorMessage,
	}
	response := Response{
		Meta: meta,
		Data: nil,
	}
	ctx.JSON(status, response)
}
