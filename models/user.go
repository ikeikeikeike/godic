package models

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/martini-contrib/sessionauth"
)

// User can be any struct that represents a user in my system
type User struct {
	Model

	Email    string    `sql:"type:varchar(128);not null" form:"email"  binding:"required"`
	Password string    `sql:"type:varchar(128);not null" form:"password"  binding:"required"`
	LoggedAt time.Time `form:"-" `

	authenticated bool `sql:"-" form:"-" `
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() sessionauth.User {
	return &User{}
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (m *User) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	m.LoggedAt = time.Now().Add(time.Duration(9) * time.Hour)
	m.authenticated = true

	DB.Save(&m)
}

// Logout will preform any actions that are required to completely
// logout a user.
func (m *User) Logout() {
	// Remove from logged-in user's list
	// etc ...
	m.authenticated = false
}

func (m *User) IsAuthenticated() bool {
	return m.authenticated
}

func (m *User) UniqueId() interface{} {
	return m.ID
}

// GetById will populate a user object from a database model with
// a matching id.
func (m *User) GetById(id interface{}) error {
	err := DB.First(&m, id).Error
	if err != nil {
		log.Warning(err)
		return err
	}

	return nil
}
