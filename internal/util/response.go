package util

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"OK"`
	Result  interface{} `json:"result"`
}

type Paginate struct {
	Total       int64       `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	Data        interface{} `json:"data"`
}

func NewError(ctx *gin.Context, status int, err error) {
	er := JsonResponse{
		Status:  status,
		Message: err.Error(),
		Result:  nil,
	}
	ctx.JSON(status, er)
}

func NewResponse(ctx *gin.Context, status int, message string, data interface{}) {
	result := JsonResponse{
		status,
		message,
		data,
	}
	ctx.JSON(status, result)
}
