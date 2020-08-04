package core

import (
	"context"
	"os"

	"github.com/Al-un/emprev-api/internals/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// MongoClient is the Mongo client instance
	MongoClient *mongo.Client
	// MongoDatabase is the EmpRev database instance in the MongoDB
	MongoDatabase *mongo.Database
)

func init() {
	// --- Parameters definition
	dbConnectionString := "mongodb://localhost:27017/emprev"
	dbName := "emprev"

	if dbURLVar := os.Getenv("DB_URL"); dbURLVar != "" {
		dbConnectionString = dbURLVar
	}
	if dbNameVar := os.Getenv("DB_NAME"); dbNameVar != "" {
		dbName = dbNameVar
	}

	// --- Connection
	utils.APILogger.Infof("Connecting to DB %s\n", dbConnectionString)
	clientOptions := options.Client().ApplyURI(dbConnectionString)
	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		utils.APILogger.Fatalf("MongoDB connection error to %v\n", dbConnectionString)
		return
	}

	// --- Check the connection
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		utils.APILogger.Fatalf("MongoDB ping to %v failed. Connection error\n", dbConnectionString)
	}

	// --- Init the database instance
	MongoDatabase = MongoClient.Database(dbName)
}
