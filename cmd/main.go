package main

import (
	"pizzaria/internal/data"
	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", handler.GetPizzas)
	router.POST("/pizzas", handler.PostPizzas)
	router.GET("/pizzas/:id", handler.GetPizzasByID)

	//Deletar uma pizza
	router.DELETE("/pizzas/:id", handler.DeletePizzaById)
	//Editar ou atualizar uma pizza
	router.PUT("/pizzas/:id", handler.UpdatePizzaById)
	//Excluir uma pizza

	err := router.Run()
	if err != nil {
		return
	}
}
