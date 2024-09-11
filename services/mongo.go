package services

import (
	"context"
	"log"

	"github.com/jherrma/gems/config"
	"github.com/jherrma/gems/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	client *mongo.Client
}

func NewMongoDb(client *mongo.Client) *MongoDb {
	return &MongoDb{client: client}
}

func (m *MongoDb) GetCollection(collectionName string) *mongo.Collection {
	return m.client.Database(config.DATABASE).Collection(collectionName)
}

func (m *MongoDb) Close() {
	m.client.Disconnect(context.Background())
}

func (m *MongoDb) EnsureMongoDbIsSetUp() {
	err := m.client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("could not ping mongo db!")

	}
	//database := client.Database(config.DATABASE)

	// create indicies
}

func (m *MongoDb) InsertGem(gem *models.Gem) error {
	collection := m.GetCollection(config.LATIN_PHRASE_COLLECTION)

	result, err := collection.CountDocuments(context.Background(), bson.M{"phrase": gem.Phrase})
	if err != nil {
		return err
	}

	if result > 0 {
		return nil
	}

	_, err = collection.InsertOne(context.Background(), gem)
	return err
}

func (m *MongoDb) GetGem(phrase string) (*models.Gem, error) {
	collection := m.GetCollection(config.LATIN_PHRASE_COLLECTION)

	result := models.Gem{}
	err := collection.FindOne(context.Background(), bson.M{"phrase": phrase}).Decode(&result)
	return &result, err
}

func (m *MongoDb) GetGems(skip, limit int64) ([]models.Gem, error) {
	ctx := context.Background()
	collection := m.GetCollection(config.LATIN_PHRASE_COLLECTION)
	options := options.Find().SetSkip(skip).SetLimit(limit)

	var gems []models.Gem
	cursor, err := collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var gem models.Gem
		err := cursor.Decode(&gem)
		if err != nil {
			return nil, err
		}
		gems = append(gems, gem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return gems, nil
}
