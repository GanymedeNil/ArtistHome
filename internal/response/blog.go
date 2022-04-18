package response

import (
	"github.com/GanymedeNil/GoFrameworkBase/internal/util"
)

type BlogPostsResponse struct {
	util.JsonResponse
	Result []BlogPostResult `json:"result"`
}

type BlogPostResult struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	Summary    string   `json:"summary"`
	Content    string   `json:"content"`
	Status     uint     `json:"status"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
}

type BlogTagsResponse struct {
	util.JsonResponse
	Result []BlogTagResult `json:"result"`
}

type BlogTagResult struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BlogCategoriesResponse struct {
	util.JsonResponse
	Result []BlogTagResult `json:"result"`
}

type BlogCategoryResult struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
