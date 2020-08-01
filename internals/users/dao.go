package users

import (
	"context"

	"github.com/Al-un/emprev-api/internals/core"
	"github.com/Al-un/emprev-api/internals/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbUserCollectionName string
	dbUserCollection     *mongo.Collection
	dbSuperAdminUserName string
)

func init() {
	dbUserCollectionName = "emprev_users"
	dbUserCollection = core.MongoDatabase.Collection(dbUserCollectionName)

	// ensure superadmin is always defined
	dbSuperAdminUserName = "root"
	filter := bson.M{
		"username": dbSuperAdminUserName,
	}
	var superAdmin core.User
	if err := dbUserCollection.FindOne(context.TODO(), filter).Decode(&superAdmin); err != nil {
		superAdmin := userWithPassword{
			User: core.User{
				ID:        primitive.NewObjectID(),
				IsAdmin:   true,
				IsDeleted: false,
				Username:  dbSuperAdminUserName,
			},
			Password: dbSuperAdminUserName,
		}
		insert, err := createUser(superAdmin)

		if err == nil {
			utils.ApiLogger.Infof("Successfully created super admin user with ID %v\n", insert.ID)
		} else {
			utils.ApiLogger.Fatalf("Error when creating super admin user, application cannot continue\n")
		}
	} else {
		utils.ApiLogger.Infof("Super admin user <%s> already exists\n", dbSuperAdminUserName)
	}
}

func createUser(user userWithPassword) (*core.User, error) {
	hashedPassword := hashPassword(user.Password)

	user.ID = primitive.NewObjectID()
	user.IsDeleted = false
	user.Password = hashedPassword

	createdUser, err := dbUserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	// Fetch the user
	var newUser core.User
	filter := bson.M{"_id": createdUser.InsertedID}
	if err := dbUserCollection.FindOne(context.TODO(), filter).Decode(&newUser); err != nil {
		return nil, err
	}

	return &newUser, nil
}

func deleteUser(userID string) (int64, error) {
	id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": id}
	d, err := dbUserCollection.DeleteMany(context.TODO(), filter, nil)
	if err != nil {
		return -1, nil
	}

	return d.DeletedCount, nil
}

func findActiveUsernamePassword(username, password string) (*core.User, error) {
	hashedPassword := hashPassword(password)

	var user userWithPassword

	filter := bson.M{
		"username":  username,
		"password":  hashedPassword,
		"isDeleted": false,
	}

	if err := dbUserCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user.User, nil
}

func getUser(userID string) (*core.User, error) {
	id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": id}

	var user core.User

	if err := dbUserCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func listUsers() (*[]core.User, error) {
	users := make([]core.User, 0)

	cur, err := dbUserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var next core.User
	for cur.Next(context.TODO()) {
		cur.Decode(&next)
		users = append(users, next)
	}

	return &users, nil
}

func updateUser(userID string, user core.User) (core.User, error) {
	id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": id}

	var returnOpt options.ReturnDocument = 1

	options := &options.FindOneAndUpdateOptions{
		ReturnDocument: &(returnOpt),
	}

	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
			"isAdmin":  user.IsAdmin,
		},
	}

	var updatedUser core.User

	if err := dbUserCollection.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&updatedUser); err != nil {
		return core.User{}, err
	}
	return updatedUser, nil
}
