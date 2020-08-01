package reviews

import (
	"context"

	"github.com/Al-un/emprev-api/internals/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbReviewCollectionName string
	dbReviewCollection     *mongo.Collection
)

func init() {
	dbReviewCollectionName = "emprev_reviews"
	dbReviewCollection = core.MongoDatabase.Collection(dbReviewCollectionName)
}

func createReview(review Review) (*Review, error) {
	review.ID = primitive.NewObjectID()

	created, err := dbReviewCollection.InsertOne(context.TODO(), review)
	if err != nil {
		return nil, err
	}

	var newReview Review
	filter := bson.M{"_id": created.InsertedID}
	if err := dbReviewCollection.FindOne(context.TODO(), filter).Decode(&newReview); err != nil {
		return nil, err
	}

	return &newReview, nil
}

func listReviews(reviewerUserID string) (*[]Review, error) {
	list := make([]Review, 0)

	filter := bson.M{
		"reviewerUserId": reviewerUserID,
	}
	cur, err := dbReviewCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var next Review
	for cur.Next(context.TODO()) {
		cur.Decode(&next)
		list = append(list, next)
	}

	return &list, nil
}

func updateReview(reviewID string, review Review) (*Review, error) {
	id, _ := primitive.ObjectIDFromHex(reviewID)
	filter := bson.M{"_id": id}

	var returnOpt options.ReturnDocument = 1

	options := &options.FindOneAndUpdateOptions{
		ReturnDocument: &(returnOpt),
	}

	update := bson.M{
		"$set": bson.M{
			"score":   review.Score,
			"comment": review.Comment,
		},
	}

	var updated Review

	if err := dbReviewCollection.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}
