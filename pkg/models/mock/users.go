package mock

import (
	"time"

	"github.com/cabralfbenja/snippetbox/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Alice",
	Email:   "alice@gmail.com",
	Created: time.Now(),
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "alice@gmail.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
