package ApiwithGoChi

import "time"

type Person struct {
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Address     string `json:"address" bson:"address"`
	Company     string `json:"company" bson:"company"`
	CreatedOn    time.Time `json:"create_on" bson:"create_on"`
}
