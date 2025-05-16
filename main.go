package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Operacao struct {
	ID        string  `json:"id"`
	Operando1 float64 `json:"operando1"`
	Operando2 float64 `json:"operando2"`
	Operacao  string  `json:"operacao"`
	Resultado float64 `json:"resultadoe"`
}

// slice
var operacoes = []Operacao{
	{ID: "1", Operando1: 10.0, Operando2: 5.0, Operacao: "somar", Resultado: 15.0},
	{ID: "2", Operando1: 10.0, Operando2: 5.0, Operacao: "subtrair", Resultado: 5.0},
	{ID: "3", Operando1: 10.0, Operando2: 5.0, Operacao: "multiplicar", Resultado: 50.0},
	{ID: "1", Operando1: 10.0, Operando2: 5.0, Operacao: "dividir", Resultado: 2.0},
}

// rotas http
func main() {
	router := gin.Default()
	router.GET("/operacoes", getOperacoes)
	router.GET("/operacoess/:id", getOperacaoID)
	router.POST("/operacoes", postOperacoes)

	router.Run("localhost:8080")
}

func getOperacoes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, operacoes)
}

func postOperacoes(c *gin.Context) {
	var novaOperacao Operacao

	if err := c.BindJSON(&novaOperacao); err != nil {
		return
	}

	operacoes = append(operacoes, novaOperacao)
	c.IndentedJSON(http.StatusCreated, novaOperacao)
}

func getOperacaoID(c *gin.Context) {
	id := c.Param("id")
	//c é o contexto da requisição HTTP (do tipo *gin.Context).
	//Param("id") é um método que pega o valor de um parâmetro da URL, chamado "id".

	for _, a := range operacoes {
		if a.ID == id {
			// Só verifica divisão por zero se for uma operação de divisão
			if a.Operacao == "divisao" && a.Operando2 == 0 {
				http.Error(c.Writer, "Erro: nenhum número pode ser dividido por zero", http.StatusBadRequest)
				return
			}

			// Retorna operação normalmente
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "operacao invalida"})
}
