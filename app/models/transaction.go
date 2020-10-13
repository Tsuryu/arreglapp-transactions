package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Transaction : user transactions model
type Transaction struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"-"`
	TraceID      string              `bson:"trace_id" json:"trace_id,omitempty"`
	SecurityCode string              `bson:"security_code" json:"security_code,omitempty"`
	Reference    string              `json:"reference"`
	Active       bool                `json:"-"`
	Email        string              `json:"email"`
	Details      []TransactionDetail `json:"details,omitempty"`
}
