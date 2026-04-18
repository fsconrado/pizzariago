package handler

import "C"
import (
	"fmt"
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatePizzaById(c *gin.Context) {
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

	if err := service.ValidatePizzaPrice(&updatePizza); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("UPDATEPIZZA: ", updatePizza)

	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			data.Pizzas[i] = updatePizza
			data.Pizzas[i].ID = id
			data.SavePizza()
			c.JSON(http.StatusOK, data.Pizzas[i])
			return
		}

	}
	c.JSON(http.StatusNotFound, gin.H{"message": "pizza não encontrada"})
}

func DeletePizzaById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()
			c.JSON(http.StatusOK, gin.H{"message": "Pizza deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})

}

func GetPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, error := strconv.Atoi(idParam)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	for _, pizza := range data.Pizzas {
		if pizza.ID == id {
			c.JSON(http.StatusOK, pizza)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
	c.JSON(http.StatusCreated, newPizza)
}

func GetPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pizzas": data.Pizzas,
	})
}
