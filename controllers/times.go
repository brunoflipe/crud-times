package controllers

import "github.com/gin-gonic/gin"

func ListarTimes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para listar times",
	})
}

func CriarTime(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para criar time",
	})
}
func AtualizarTime(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para atualizar time",
	})
}
func DeletarTime(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para deletar time",
	})
}
func LerTime(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mensagem": "rota para ler time",
	})
}
