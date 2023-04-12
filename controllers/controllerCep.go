package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	structModel "github.com/guilherm5/cep/struct"
)

func ObtemCep(c *gin.Context) {
	var getCep structModel.CEP

	if err := c.ShouldBindJSON(&getCep); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro ao decodificar body para receber solicitação": err,
		})
		log.Println("Erro ao decodificar body para receber solicitação", err)
		return
	}

	if len(getCep.CEP) > 9 || len(getCep.CEP) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro ao Realizar solicitação": "CEP Digitado deve conter 8 ou 9 (contando com -) caracteres.",
		})
		log.Println("Erro, CEP Digitado deve conter 8 ou 9 (contando com -) caracteres.")
	} else if len(getCep.CEP) == 8 || len(getCep.CEP) == 9 {

		resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", getCep.CEP))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao obter CEP": err,
			})
			log.Println("Erro ao obter CEP", err)
			return
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&getCep); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao receber resposta": err,
			})
			log.Println("Erro ao receber resposta", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"cep":         getCep.CEP,
			"logradouro":  getCep.Logradouro,
			"complemento": getCep.Complemento,
			"bairro":      getCep.Bairro,
			"estado":      getCep.Localidade,
			"uf":          getCep.UF,
		})
	}
}
