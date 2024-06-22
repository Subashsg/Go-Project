package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Developer struct {
	ID       primitive.ObjectID
	Name     string `json:"name"`
	Age      int    `json:"age"`
	E_mail   string `json:"e_mail"`
	Contact  int    `json:"contact"`
	Position string `json:"position"`
}
