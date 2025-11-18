package mongodb

import (
	"time"
)

type User struct {
	ID        string    `bson:"userId"`
	Username  string    `bson:"username"`
	Email     string    `bson:"email"`
	Password  string    `bson:"passwordHash"`
	Role      string    `bson:"role"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	Version   int32     `bson:"version"`
}
