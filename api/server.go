package api

import (
	db "github.com/ashiqur/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// server serve the HTTP request for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/accounts", server.listAccounts)
	router.GET("/account/:id", server.getAccount)
	router.POST("/account", server.createAccount)
	router.PUT("/account/update", server.updateAccount)
	router.DELETE("/account/delete/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on specific port
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
