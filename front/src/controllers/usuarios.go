package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("nome"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(usuario)
}
