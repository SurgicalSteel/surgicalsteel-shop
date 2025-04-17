package server

import (
	"net/http"
	"surgicalsteel-shop/controller"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func BuildHttpServer() *http.Server {

	httpServer := &http.Server{
		Addr:         ":9090",
		Handler:      buildRoutingHandler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return httpServer
}

func buildRoutingHandler() *gin.Engine {
	engine := gin.Default()
	// if config.GINMode == resources.Release {
	// 	engine = gin.New()
	// } else {
	// 	engine = gin.Default()
	// }

	noAuth := engine.Group("/api")
	{
		noAuth.GET("/hello", controller.HandleHello)

		noAuth.GET("/buyer/quotes/:id", controller.HandleGetBuyerQuoteById)
	}

	engine.Use(gzip.Gzip(gzip.BestCompression))
	return engine
}
