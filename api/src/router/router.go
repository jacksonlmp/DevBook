package router

import "github.com/gorilla/mux"

// Retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}