package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"	
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// PACOTE MIDDLEWARES FICA ENTRE A REQUISÃO E A RESPOSTA

// é a mesma função "faz a mesma coisa" func (w http.ResponseWriter, r *http.Request)
// Autenticar verifica se o usuário fazendo a requisção está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return			
		}
		proximaFuncao(w, r)
	}
}