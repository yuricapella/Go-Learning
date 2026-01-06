package user

import (
	"time"

	"github.com/google/uuid"
)

// NewUser cria uma nova instância de User com ID uuid e timestamps preenchidos.
func NewUser(name string) User {
	return User{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
