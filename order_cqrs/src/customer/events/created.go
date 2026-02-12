package events

import "time"

type Created struct {
	ID        int64     `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"created_at" bson:"createdat"`
}
