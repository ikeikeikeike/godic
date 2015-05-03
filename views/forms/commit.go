package forms

type Commit struct {
	Name    string `form:"name" binding:"required"`
	Yomi    string `form:"yomi" binding:"required"`
	Content string `form:"content" binding:"required"`
	Outline string `form:"outline"`
	Message string `form:"message"`
}

// func (p Post) Validate(errors binding.Errors, req *http.Request) binding.Errors { return errros }
