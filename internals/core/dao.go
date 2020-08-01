package core

import (
	"context"

	"github.com/Al-un/emprev-api/internals/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient   *mongo.Client
	MongoDatabase *mongo.Database
)

func init() {
	dbConnectionString := "mongodb://localhost:27017"
	dbName := "emprev"

	clientOptions := options.Client().ApplyURI(dbConnectionString)
	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		utils.ApiLogger.Fatalf("MongoDB connection error to %v\n", dbConnectionString)
		return
	}

	// Check the connection
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		utils.ApiLogger.Fatalf("MongoDB ping to %v failed. Connection error\n", dbConnectionString)
	}

	// Init the database instance
	MongoDatabase = MongoClient.Database(dbName)
}
