package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// gerar router das rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
