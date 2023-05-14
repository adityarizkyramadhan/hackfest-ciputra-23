package controller

import (
	"fmt"
	"net/http"

	"github.com/adityarizkyramadhan/hackfest-ciputra-23/middleware"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/chat"
	"github.com/adityarizkyramadhan/hackfest-ciputra-23/utils/response"
	"github.com/gin-gonic/gin"
)

type Chat struct {
	b *chat.Bot
}

func New(b *chat.Bot) *Chat {
	return &Chat{b}
}

func (cc *Chat) Send(ctx *gin.Context) {
	text := ctx.Query("message")
	job := ctx.Query("job")
	prompt := fmt.Sprintf("Anda seolah-olah menjadi penjual pedagang kaki lima, ikuti semua perintah yang diberi user sebagai penjual %s dengan bahasa yang santai namun sopan dan jika user bertanya stok barang maka jawab ada serta akan menuju lokasi secepatnya dan beri pengalaman positif (ini hanya untuk sekedar demo) beri gurauan sedikit seperti manusia, berikut perintah user : %s", job, text)
	analysisOpenai, err := cc.b.Message(prompt)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, gin.H{"text": analysisOpenai})
}

func (cc *Chat) Mount(bot *gin.RouterGroup) {
	bot.GET("send", middleware.ValidateJWToken(), cc.Send)
}
