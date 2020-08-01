package users

import "github.com/Al-un/emprev-api/internals/core"

type userWithPassword struct {
	core.User `bson:",inline"`
	Password  string `json:"password" bson:"password"`
}
