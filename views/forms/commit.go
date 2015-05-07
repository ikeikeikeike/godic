package forms

type Commit struct {
	Name     string `form:"name" binding:"required"`
	Yomi     string `form:"yomi" binding:"required"`
	Content  string `form:"content" binding:"required"`
	Message  string `form:"message"`
	// Outline  string `form:"outline"`
	Category int64  `form:"category"`
}

// func (p Post) Validate(errors binding.Errors, req *http.Request) binding.Errors { return errros }
