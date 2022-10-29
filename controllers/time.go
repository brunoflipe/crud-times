package controllers

import (
	"strconv"

	"github.com/brunoflipe/crud-times/models"
	"github.com/gin-gonic/gin"
)

func ListarTimes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"times": models.ListaDeTimes,
	})
}

func LerTime(ctx *gin.Context) {
	// receber o id da requisição
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "id é obrigatorio",
		})
		return
	}

	newid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Id tem q ser inteiro",
		})

		return
	}
	// achar o id correto
	var time *models.Time
	for i := 0; i < len(models.ListaDeTimes); i++ {
		if newid == models.ListaDeTimes[i].ID {
			time = models.ListaDeTimes[i]
			break
		}
	}

	if time == nil {
		ctx.JSON(404, gin.H{
			"mensagem": "time nâo encontrado",
		})
		return
	}

	// se achou o id, então retorna o time encontrado

	ctx.JSON(200, gin.H{
		"time": time,
	})
}

func CriarTime(ctx *gin.Context) {
	// pegar o time do JSON e converter para struct do Go
	var time *models.Time

	err := ctx.BindJSON(&time)

	// se deu erro na conversao, entao imprime o erro e nao continua
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// verificar se o id do time que está sendo enviado no JSON já existe no array

	// adicionar essa struct no array
	models.ListaDeTimes = append(models.ListaDeTimes, time)

	// imprimir mensagem de criado com sucesso e os dados time inserido
	ctx.JSON(200, gin.H{
		"mensagem": "time criado com sucesso",
		"time":     time,
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
