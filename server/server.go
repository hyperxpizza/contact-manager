package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	database "github.com/hyperxpizza/contact-manager/server/database"
	"github.com/hyperxpizza/contact-manager/server/graph"
	"github.com/hyperxpizza/contact-manager/server/graph/generated"
	"github.com/hyperxpizza/contact-manager/server/handlers"
	"github.com/joho/godotenv"
)

const defaultPort = 8080

func main() {
	router := gin.Default()

	err := godotenv.Load("./../.mongo.env")
	if err != nil {
		log.Printf("[-] Loading enviroment files failed: %v", err)
	}

	mongoUser := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongoPassword := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	mongoHost := os.Getenv("MONGO_HOST")
	mongoAuthSource := os.Getenv("MONGO_AUTH_SOURCE")
	mongoPort := os.Getenv("MONGO_PORT")

	port, err := strconv.Atoi(mongoPort)
	if err != nil {
		log.Printf("[-] Can not convert mongoPort to integer: %v\n", err)
	}

	database.ConnectToDB(mongoUser, mongoPassword, mongoAuthSource, mongoHost, port)

	router.Use(cors.Default())
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.GET("/download/:bucket/*filename", handlers.DownloadFile)

	log.Println("[*] Server running at: " + strconv.FormatInt(defaultPort, 10))
	router.Run(fmt.Sprintf(":%d", defaultPort))
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
