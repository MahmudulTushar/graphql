package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MahmudulTushar/graphql/graph"
	"github.com/MahmudulTushar/graphql/graph/generated"
	"github.com/gin-gonic/gin"
)

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}

func GraphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(context *gin.Context) {
		srv.ServeHTTP(context.Writer, context.Request)
	}
}
