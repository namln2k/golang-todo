package v1Handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (u *TodoHandler) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": []string{}, "message": "Get all todos v1"})
}
