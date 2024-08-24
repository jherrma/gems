package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/jherrma/gems/config"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE                = "gems"
	LATIN_PHRASE_COLLECTION = "latin-phrases"
)

func closeMongoConnection(client *mongo.Client) {
	if err := client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func ensureMongoDbIsSetUp(ctx context.Context, client *mongo.Client) {
	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("could not ping mongo db!")
	}
	//database := client.Database(DATABASE)

	// create indicies
}

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
	defer closeMongoConnection(client)
	ensureMongoDbIsSetUp(ctx, client)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := os.Getenv(config.SERVER_PORT)
	if port == "" {
		log.Fatal("PORT for server must be set!")
	}

	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	app.Listen(port)
}
