package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/jherrma/gems/config"
	"github.com/jherrma/gems/handlers"
	"github.com/jherrma/gems/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	_ = godotenv.Load()

	uri := os.Getenv(config.MONGODB_URI)
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. See: www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	mongoDb := services.NewMongoDb(client)
	defer mongoDb.Close()
	mongoDb.EnsureMongoDbIsSetUp()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("")
	})

	app.Post("/api/gem", handlers.InsertGem(mongoDb))
	app.Post("/api/phrase", handlers.InsertPhrase(mongoDb))
	app.Post("/api/nearestItems", handlers.GetNearesItems(mongoDb))
	app.Get("/api/list", handlers.GetList(mongoDb))

	port := os.Getenv(config.SERVER_PORT)
	if port == "" {
		log.Fatal("PORT for server must be set!")
	}

	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	app.Listen(port)
}
