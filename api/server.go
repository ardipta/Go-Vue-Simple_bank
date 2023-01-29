package api

import (
	"fmt"

	"github.com/ashiqur/simplebank/token"
	"github.com/ashiqur/simplebank/util"

	db "github.com/ashiqur/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// server serve the HTTP request for our banking service
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	// For JWT TOKEN => NewJWTMaker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can not create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	router.GET("/accounts", server.listAccounts)
	router.GET("/account/:id", server.getAccount)
	router.POST("/account", server.createAccount)
	router.PUT("/account/update", server.updateAccount)
	router.DELETE("/account/delete/:id", server.deleteAccount)

	router.POST("/transfers", server.createTransfer)
	server.router = router
}

// Start runs the HTTP server on specific port
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
