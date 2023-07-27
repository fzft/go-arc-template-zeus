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
	*gin.Engine
	listenAddr string
}

func New() *Server {
	r := gin.Default()
	r.Routes()
	return &Server{Engine: r}
}

// Start runs the http server. this is a blocking call. outsider caller should run it in a goroutine.
func (s *Server) Start() {
	s.listenAddr = fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))

	if err := s.Run(s.listenAddr); err != nil {
		Logger.Error(fmt.Sprintf("server run error: %v", err))
	}
}

func (s *Server) Close() {
	s.S
}

// Routes sets the routes for the http server.
func Routes(r *gin.Engine) {
	rg := r.Group("/api")

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
