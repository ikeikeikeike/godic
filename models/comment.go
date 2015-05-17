package models

import (
	"database/sql"

	"github.com/ikeikeikeike/godic/modules/funcmaps"
)

type Comment struct {
	Model

	Title       string `sql:"type:varchar(64)" form:"title"`                        // gin index
	Comment     string `sql:"type:text;not null" form:"comment" binding:"required"` // gin index
	CommentHTML string `sql:"type:text;not null" form:"-"`

	Role string `sql:"type:varchar(32);index;default:'public'" form:"-"`

	ObjectID   int64  `sql:"index" form:"-"`
	ObjectType string `sql:"index" form:"-"`

	User   *User
	UserID sql.NullInt64 `sql:"index" form:"-"`
}

func (m *Comment) BeforeSave() error {
	m.CommentHTML = funcmaps.AutoLink(funcmaps.MarkdownHTML(m.Comment), cachedDictNames())
	return nil
}
