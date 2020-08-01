package reviews

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	ReviewerID primitive.ObjectID `json:"reviewerUserId,omitempty" bson:"reviewerUserId,omitempty"`
	ReviewedID primitive.ObjectID `json:"reviewedUserId,omitempty" bson:"reviewedUserId,omitempty"`
	Period     string             `json:"period,omitempty" bson:"period,omitempty"`
	Score      int                `json:"score,omitempty" bson:"score,omitempty"`
	Comment    string             `json:"comment,omitempty" bson:"comment,omitempty"`
}
