package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guilherm5/cep/database"

	structModel "github.com/guilherm5/cep/struct"
)

var DB = database.Init()
var getCep structModel.CEP

func ObtemCep(c *gin.Context) {

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
		result, err := DB.Exec(`INSERT INTO cep (cep, logradouro, complemento, bairro, localidade, uf) VALUES ($1, $2, $3, $4, $5, $6)`, getCep.CEP, getCep.Logradouro, getCep.Complemento, getCep.Bairro, getCep.Localidade, getCep.UF)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro ao inserir dados na tabela ": err,
			})
			log.Println("Erro ao inserir dados na tabela ", err)
			return
		}
		log.Println("Sucesso ao realizar insert na tabela CEP", result)
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
