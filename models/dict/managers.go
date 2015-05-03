package dict

import (
	m "github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/views/forms"
)

func FirstOrCreateByCommit(c forms.Commit) (*m.Dict, bool) {
	var d m.Dict
	m.DB.Where(m.Dict{Name: c.Name}).FirstOrInit(&d)

	// created
	created := false
	if d.ID <= 0 {
		created = true

		d.Kana = c.Kana
		d.Outline = c.Outline

		// if v, ok := p["image"]; ok {
		// d.Image = &m.Image{}
		// _ = v
		// }

		// if v, ok := p["category"]; ok {
		// d.Category = &m.Category{}
		// _ = v
		// }

		m.DB.Save(&d)
	}

	return &d, created
}
