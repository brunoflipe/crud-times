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
		n1.PUT("/atualizar", atualizarTimes)
		n1.DELETE("/deletar", deleteTimes)
	}
	n2 := r.Group("/tecnico")
	{
		n2.GET("/listar", listarTecnicos)
		n2.POST("/criar", criarTecnicos)
		n2.PUT("/atualizar", atualizarTecnicos)
		n2.DELETE("/deletar", deleteTecnicos)
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
func atualizarTimes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para atualizar times",
	})
}
func deleteTimes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para deletar times",
	})
}
func listarTecnicos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para listar tecnico",
	})
}

func criarTecnicos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para criar tecnico",
	})
}
func atualizarTecnicos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para atualizar tecnico",
	})
}
func deleteTecnicos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para deletar tecnico",
	})
}
