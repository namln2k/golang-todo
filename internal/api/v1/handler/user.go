package v1Handler

import (
	"example.com/golang-todo/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSerive service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userSerive: service,
	}
}

func (handler *UserHandler) GetAll(ctx *gin.Context) {
	handler.userSerive.GetAll()
}

func (handler *UserHandler) GetOneById(ctx *gin.Context) {
}

func (handler *UserHandler) Create(ctx *gin.Context) {
}

func (handler *UserHandler) Update(ctx *gin.Context) {
}

func (handler *UserHandler) Delete(ctx *gin.Context) {
}
