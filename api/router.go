package api

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/asadbekGo/api-gateway/api/handlers/v1"
	"github.com/asadbekGo/api-gateway/config"
	"github.com/asadbekGo/api-gateway/pkg/logger"
	"github.com/asadbekGo/api-gateway/services"
)

// Option ...

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/todos", handlerV1.CreateTodo)
	api.GET("/todos/:id", handlerV1.GetTodo)
	api.GET("/todos", handlerV1.ListTodos)
	api.PUT("/todos/:id", handlerV1.UpdateTodo)
	api.DELETE("/todos/:id", handlerV1.DeleteTodo)
	api.GET("/todoListOverdue", handlerV1.ListOverdueTodo)

	return router
}
