package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TransactionDetail : transaction historic
type TransactionDetail struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Status   string             `json:"status"`
	Date     time.Time          `json:"date"`
	Metadata interface{}        `bson:"metadata,omitempty" json:"metadata,omitempty"`
}
