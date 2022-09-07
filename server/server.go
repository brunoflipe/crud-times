package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateServer() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	n1 := r.Group("/times")
	{
		n1.GET("/listar", listarTimes)
		n1.POST("/criar", criarTimes)

	}

	fmt.Println("API escutando na porta 2525")
	err := r.Run(":2525")
	if err != nil {
		log.Fatalf("nao foi possivel rodar a api na porta 2525: %v", err)
	}
}

func listarTimes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para listar times",
	})
}

func criarTimes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para criar times",
	})
}
