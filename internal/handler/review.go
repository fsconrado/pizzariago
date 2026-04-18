package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	var newReview models.Review

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidatePizzaReview(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			data.Pizzas[i].Review = append(data.Pizzas[i].Review, newReview)
			data.SavePizza()
			c.JSON(http.StatusCreated, data.Pizzas[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})

}
