package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type IServer interface {
}

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	return &Server{engine: gin.New()}
}

// Run runs the http server. this is a blocking call. outsider caller should run it in a goroutine.
func (s *Server) Run() {
	s.Routes()
	listenAddr := fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))

	// blocking call
	s.engine.Run(listenAddr)
}

// Routes sets the routes for the http server.
func (s *Server) Routes() {
	rg := s.engine.Group("/api")

	addUserRoutes(rg)
	addPingRoutes(rg)
}

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
