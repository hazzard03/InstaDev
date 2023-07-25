package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                  "/publicacoes",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarPublicacao,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublicacoes,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicacaoId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublicacao,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicacaoId}",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarPublicacao,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicacaoId}",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarPublicacao,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/usuarios/{usuarioId}/publicacoes",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublicacoesPorUsuario,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicacaoId}/curtir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CurtirPublicacao,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicacaoId}/descurtir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.DescurtirPublicacao,
		RequerAutentificacao: true,
	},
}
