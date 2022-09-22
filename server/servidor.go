package server

import (
	"fmt"
	"log"

	"github.com/brunoflipe/crud-times/models"
	"github.com/brunoflipe/crud-times/routes"
	"github.com/gin-gonic/gin"
)

func CreateServer() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	routes.AddRoutes(r)

	models.ListaDeTimes = make([]*models.Time, 0)

	fmt.Println("API escutando na porta 2525")
	err := r.Run(":2525")
	if err != nil {
		log.Fatalf("nao foi possivel rodar a api na porta 2525: %v", err)
	}
}
