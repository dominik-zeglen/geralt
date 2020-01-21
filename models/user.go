package models

import "fmt"

const UsersCollectionKey = "users"

type User struct {
	BaseModel `bson:",inline"`
	Email     string
	Name      string
}

func (u User) String() string {
	return fmt.Sprintf("User<%s>", u.ID)
}
