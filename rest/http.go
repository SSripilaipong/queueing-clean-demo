package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

func newHttpServer(port string, router *gin.Engine, c *cors.Cors, readTimeout time.Duration, writeTimeout time.Duration) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      c.Handler(makeHandler(router)),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}

func newCorsHandler() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	return c
}

func makeHandler(routes http.Handler) *negroni.Negroni {
	handler := negroni.New(negroni.NewRecovery())
	handler.UseHandler(routes)
	return handler
}
