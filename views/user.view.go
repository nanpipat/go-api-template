package views

import "github.com/nanpipat/go-api-template/models"

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewUser(m *models.User) *User {
	return &User{
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
}
