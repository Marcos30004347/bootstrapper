package network

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RESTServer struct {
	server *http.Server
}

var rest_server *RESTServer = nil

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", ping)

	return e
}

func RunRESTServer() error {
	return rest_server.server.ListenAndServe()
}

func InitRESTServer(port int) {
	rest_server = &RESTServer{}

	rest_server.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
