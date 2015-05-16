package models

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/martini-contrib/sessionauth"
)

type User struct {
	Model

	Email    string    `sql:"type:varchar(128);not null;unique" form:"email"  binding:"required"`
	Password string    `sql:"type:varchar(128);not null" form:"password"  binding:"required"`
	LoggedAt time.Time `form:"-" `

	Repassword    string `sql:"-" form:"repassword"`
	authenticated bool   `sql:"-" form:"-" `
}

func GenerateAnonymousUser() sessionauth.User {
	return &User{}
}

func (m *User) Login() {
	m.LoggedAt = time.Now().Add(time.Duration(9) * time.Hour)
	m.authenticated = true

	DB.Save(&m)
}

func (m *User) Logout() {
	m.authenticated = false
}

func (m *User) IsAuthenticated() bool {
	return m.authenticated
}

func (m *User) UniqueId() interface{} {
	return m.ID
}

func (m *User) GetById(id interface{}) error {
	err := DB.First(&m, id).Error
	if err != nil {
		log.Warning(err)
		return err
	}

	return nil
}
