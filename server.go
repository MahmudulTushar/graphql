package main

import (
	http2 "github.com/MahmudulTushar/graphql/http"
	"github.com/MahmudulTushar/graphql/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "1111"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.Default()
	server.Use(middleware.Auth())
	server.GET("/", http2.PlaygroundHandler())
	server.POST("/query", http2.GraphqlHandler())
	server.Run()
}
