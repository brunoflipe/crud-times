package routes

import (
	"github.com/brunoflipe/crud-times/controllers"
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	n1 := r.Group("/times")
	{
		n1.GET("/listar", controllers.ListarTimes)
		n1.GET("/listar/:id", controllers.LerTime)
		n1.POST("/criar", controllers.CriarTime)
		n1.PUT("/atualizar", controllers.AtualizarTime)
		n1.DELETE("/deletar", controllers.DeletarTime)
	}
}
