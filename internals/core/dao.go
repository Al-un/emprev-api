package core

import (
	"context"
	"os"

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

	if dbURLVar := os.Getenv("DB_URL"); dbURLVar != "" {
		dbConnectionString = dbURLVar
	}
	if dbNameVar := os.Getenv("DB_NAME"); dbNameVar != "" {
		dbName = dbNameVar
	}

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
