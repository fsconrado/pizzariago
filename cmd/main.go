package main

import (
	"fmt"
	"net/http"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	loadPizzas()
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasByID)

	//Deletar uma pizza
	router.DELETE("/pizzas/:id", deletePizzaById)
	//Editar ou atualizar uma pizza
	router.PUT("/pizzas/:id", updatePizzaById)
	//Excluir uma pizza

	err := router.Run()
	if err != nil {
		return
	}
}

func updatePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatePizza models.Pizza
	if err := c.ShouldBindJSON(&updatePizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("UPDATEPIZZA: ", updatePizza)

	for i, pizza := range pizzas {
		if pizza.ID == id {
			pizzas[i] = updatePizza
			pizzas[i].ID = id
			savePizza()
			c.JSON(http.StatusOK, pizzas[i])
			return
		}

	}
	c.JSON(http.StatusNotFound, gin.H{"message": "pizza não encontrada"})
}

func deletePizzaById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	for i, pizza := range pizzas {
		if pizza.ID == id {
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			savePizza()
			c.JSON(http.StatusOK, gin.H{"message": "Pizza deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})

}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, error := strconv.Atoi(idParam)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(http.StatusOK, pizza)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(http.StatusCreated, newPizza)
}

func getPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pizzas": pizzas,
	})
}
