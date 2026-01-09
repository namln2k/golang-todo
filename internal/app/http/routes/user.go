package routes

import (
	v1Handler "example.com/golang-todo/internal/api/v1/handler"
	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userHandler *v1Handler.UserHandler
}

func NewUserRoute(handler *v1Handler.UserHandler) *UserRoute {
	return &UserRoute{
		userHandler: handler,
	}
}

func (userRouter *UserRoute) Register(r *gin.RouterGroup) {
	users := r.Group("/users")

	users.GET("", userRouter.userHandler.GetAll)
	users.GET("/:uuid", userRouter.userHandler.GetOneById)
	users.POST("", userRouter.userHandler.Create)
	users.PUT("/:uuid", userRouter.userHandler.Update)
	users.DELETE("/:uuid", userRouter.userHandler.Delete)
}
