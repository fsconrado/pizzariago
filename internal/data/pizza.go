package data

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("dados/pizzas.json")
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Erro ao abrir pizza de dados: ", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Pizzas); err != nil {
		//log.Fatal(err)
		fmt.Println("Erro decode pizzas: ", err)
	}
}

func SavePizza() {
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Erro ao abrir pizza de dados: ", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Pizzas); err != nil {
		fmt.Println("Erro encode pizzas: ", err)
	}

}
