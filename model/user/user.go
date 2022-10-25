package user

import (
	"time"
)

type User struct {
	Id            string    `bson:"_id,omitempty"`
	FirstName     string    `bson:"firstName,omitempty"`
	Age           int       `bson:"Age,omitempty"`
	LastName      string    `bson:"lastName,omitempty"`
	Email         string    `bson:"Email,omitempty"`
	BirthLocation string    `bson:"BirthLocation,omitempty"`
	BirthDate     string    `bson:"BirthDate,omitempty"`
	PhoneNumber   string    `bson:"PhoneNumber,omitempty"`
	UserName      string    `bson:"UserName,omitempty"`
	Password      string    `bson:"Password,omitempty"`
	RegisterDate  time.Time `bson:"RegisterDate,omitempty"`
	CreatorUserId string    `bson:"CreatorUserId,omitempty"`
	Roles         []string  `bson:"Roles,omitempty"`
}
