package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"front/src/config"
	"front/src/requisicoes"
	"front/src/respostas"
	"log" // Importe a biblioteca "log" aqui
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Printf("Iniciando a função CriarPublicacao") // Log de início da função

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		log.Printf("Erro ao criar JSON de publicação: %v", erro) // Log de erro
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		log.Printf("Erro na requisição à API: %v", erro) // Log de erro
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		log.Printf("Resposta com código de erro: %d", response.StatusCode) // Log de status de erro
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	log.Printf("Publicação criada com sucesso") // Log de sucesso
	respostas.JSON(w, response.StatusCode, nil)
}
