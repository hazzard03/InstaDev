package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                  "/usuarios",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarUsuario,
		RequerAutentificacao: false,
	},

	{
		URI:                  "/usuarios",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarUsuarios,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarUsuario,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/usuarios/{usuariosId}",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarUsuario,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarUsuario,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/usuarios/{usuarioId}/seguir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.SeguirUsuario,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.ParardeSeguirUsuario,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/usuarios/{usuarioId}/seguindo",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarSeguindo,
		RequerAutentificacao: true,
	},

	{
		URI:                  "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:               http.MethodPost,
		Funcao:               controllers.AtualizarSenha,
		RequerAutentificacao: true,
	},
}
