package users

import "github.com/Al-un/emprev-api/internals/core"

// userWithPassword covers the "password" field. Password is not
// included in the base User definition to avoid having it sent
// to the client.
type userWithPassword struct {
	core.User `bson:",inline"`
	Password  string `json:"password" bson:"password"`
}
