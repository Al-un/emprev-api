package reviews

import "go.mongodb.org/mongo-driver/bson/primitive"

// Review is the atomic element of a employee performance reivew. A
// review is done by an employee for another employee.
//
// A review is first created by an administrator (in reality, it might
// be an HR or team lead) and then the reviewer can submit the review
//
// Being atomic, any form of aggregation requires another object that
// is out of the scope for the current prototype
type Review struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	ReviewerID primitive.ObjectID `json:"reviewerUserId,omitempty" bson:"reviewerUserId,omitempty"`
	ReviewedID primitive.ObjectID `json:"reviewedUserId,omitempty" bson:"reviewedUserId,omitempty"`
	Period     string             `json:"period,omitempty" bson:"period,omitempty"`
	Score      int                `json:"score,omitempty" bson:"score,omitempty"`
	Comment    string             `json:"comment,omitempty" bson:"comment,omitempty"`
}
