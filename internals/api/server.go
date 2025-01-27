package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sulabhghimire/bank-api/internals/db/sqlc"
)

// Server will serve all banking request for http service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/accounts", server.listAccounts)

	server.router = router

	return server
}

// To run HTTP Server
func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}
