package request

//@Description title,content,status is Required;
//@Description status 1:draft; 2:pending; 3:private; 4:publish
type BlogPostForm struct {
	Title      string   `form:"title" json:"title" binding:"required" example:"Hello World"`
	Summary    string   `form:"summary" json:"summary" example:"This is first post"`
	Content    string   `form:"content" json:"content" binding:"required"  example:"{\"block\":\"something\"}"`
	Status     uint     `form:"status" json:"status" binding:"required" example:"2"`
	Tags       []string `form:"tags" json:"tags" example:"test-tag"`
	Categories []string `form:"Categories" json:"categories" example:"test-category"`
}

type BlogPostQuery struct {
	Tag      string `form:"tag"`
	Category string `form:"category"`
	Search   string `form:"search"`
	Page     int    `form:"page,default=1"`
	Limit    int    `form:"limit,default=10"`
}

//@Description name is Required;
type BlogTagForm struct {
	Name        string `form:"name" json:"name" bind:"required" example:"test-tag"`
	Description string `form:"description" json:"description" example:"this is test-tag"`
}

type BlogTagQuery struct {
	Search string `form:"search"`
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
}

//@Description name is Required;
type BlogCategoryForm struct {
	Name        string `form:"name" json:"name" bind:"required" example:"test-category"`
	Description string `form:"description" json:"description" example:"this is test-category"`
}

type BlogCategoryQuery struct {
	Search string `form:"search"`
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
}
