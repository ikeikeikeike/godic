package models

import "database/sql"

type Comment struct {
	Model

	Title   string `sql:"type:varchar(64)" form:"title"` // gin index
	Comment string `sql:"type:text" form:"comment"`      // gin index

	Role string `sql:"type:varchar(32);index;default:'public'" form:"-"`

	ObjectID   int64  `sql:"index" form:"-"`
	ObjectType string `sql:"index" form:"-"`

	User   *User
	UserID sql.NullInt64 `sql:"index" form:"-"`
}
