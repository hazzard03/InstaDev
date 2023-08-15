package main

import (
	"fmt"
	"front/src/config"
	"front/src/router"
	"front/src/utils"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
