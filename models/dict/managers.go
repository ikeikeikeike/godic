package dict

import (
	m "github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/views/forms"
	"github.com/jinzhu/gorm"
)

func RelationDB() *gorm.DB {
	return m.DB.Table("dicts").Preload("Image").Preload("Category").Preload("Comments").
		// Preload("Tags"). XXX: m2m preload does not work.
		Select("dicts.*")
}

func FirstByName(name string, out interface{}) *gorm.DB {
	return RelationDB().Where("name = ?", name).First(out)
}

func UpdateByCommit(c forms.Commit) *m.Dict {
	var d m.Dict
	FirstByName(c.Name, &d)

	do := false

	if d.Yomi != c.Yomi {
		d.Yomi = c.Yomi
		do = true
	}
	if d.Content != c.Content {
		d.Content = c.Content
		do = true
	}

	cate := &m.Category{}
	m.DB.First(cate, c.Category)
	if d.Category.ID != c.Category {
		d.Category = cate
		do = true
	}

	if do {
		m.DB.Save(&d)
	}
	return &d
}

func FirstOrCreateByCommit(c forms.Commit) (*m.Dict, bool, error) {
	var d m.Dict

	if err := m.DB.Where(m.Dict{Name: c.Name}).FirstOrInit(&d).Error; err != nil {
		return nil, false, err
	}

	// created
	created := false
	if d.ID <= 0 {
		created = true

		d.Yomi = c.Yomi
		d.Content = c.Content

		// if v, ok := p["image"]; ok {
		// d.Image = &m.Image{}
		// _ = v
		// }

		if c.Category > 0 {
			cate := &m.Category{}
			m.DB.First(cate, c.Category)
			d.Category = cate
		}

		m.DB.Save(&d)
	}

	return &d, created, nil
}
