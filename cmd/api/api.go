package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guiziin227/goApiJwt/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run monta o servidor HTTP e inicia a escuta na porta especificada.
// Ele cria um roteador principal usando o pacote Gorilla Mux, define um subroteador para a versão da API ("/api/v1") e
// registra os manipuladores de rotas do usuário.
// Por fim, ele inicia o servidor HTTP e retorna qualquer erro que ocorra durante a execução.
func (s *APIServer) Run() error {

	// O roteador principal é criado usando o pacote Gorilla Mux, q
	//ue fornece funcionalidades avançadas de roteamento para aplicativos web em Go.
	router := mux.NewRouter()

	// O subroteador é criado para agrupar as rotas da versão da API ("/api/v1").
	// Isso permite que todas as rotas relacionadas à API sejam organizadas sob um prefixo comum.
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)

	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Printf("Listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)

}
