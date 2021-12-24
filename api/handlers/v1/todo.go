package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/asadbekGo/api-gateway/api/docs" //swag
	pb "github.com/asadbekGo/api-gateway/genproto"
	l "github.com/asadbekGo/api-gateway/pkg/logger"
	"github.com/asadbekGo/api-gateway/pkg/utils"
)

// CreateTodo ...
// @Summary CreateTodo
// @Description This API for creating a new todo
// @Tags todo
// @Accept json
// @Produce json
// @Param Todo request body models.Todo true "todoCreateRequest"
// @Success 200 {object} models.Todo
// @Success 400 {object} models.StandardErrorModel
// @Success 500 {object} models.StandardErrorModel
// @Router /v1/todos/ [post]
func (h *handlerV1) CreateTodo(c *gin.Context) {
	var (
		body        pb.Todo
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	repsonse, err := h.serviceManager.TodoService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create todo", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, repsonse)
}

// GetTodo ...
// @Summary GetTodo
// @Description This API for getting todo detail
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} models.Todo
// @Success 400 {object} models.StandardErrorModel
// @Success 500 {object} models.StandardErrorModel
// @Router /v1/todos/{id} [get]
func (h *handlerV1) GetTodo(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TodoService().Get(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get todo", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListTodos ...
// @Summary ListTodos
// @Description This API for getting list of todos
// @Tags todo
// @Accept json
// @Produce json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListTodos
// @Success 400 {object} models.StandardErrorModel
// @Success 500 {object} models.StandardErrorModel
// @Router /v1/todos [get]
func (h *handlerV1) ListTodos(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TodoService().List(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list todos", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateTodo ...
// @Summary UpdateTodo
// @Description This API for updating todo
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Param Todo request body models.UpdateTodo true "todoUpdateRequest"
// @Success 200 {object} models.Todo
// @Success 400 {object} models.StandardErrorModel
// @Success 500 {object} models.StandardErrorModel
// @Router /v1/todos/{id} [put]
func (h *handlerV1) UpdateTodo(c *gin.Context) {
	var (
		body        pb.Todo
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TodoService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update todo", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTodo ...
// @Summary DeleteTodo
// @Description This API for deleting todo
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200
// @Success 400 {object} models.StandardErrorModel
// @Success 500 {object} models.StandardErrorModel
// @Router /v1/todos/{id} [delete]
func (h *handlerV1) DeleteTodo(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TodoService().Delete(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete todo", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListOverdueTodo ...
// @Summary ListOverdueTodo
// @Description This API for getting listOverdue of todo
// @Tags todo
// @Accept json
// @Produce json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Param listTime query string true "ListTime"
// @Success 200 {object} models.ListTodos
// @Success 400 {object} models.StandardErrorModel
// @Success 500 {object} models.StandardErrorModel
// @Router /v1/todoListOverdue [get]
func (h *handlerV1) ListOverdueTodo(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TodoService().ListOverdue(
		ctx, &pb.ListTime{
			ListPage: &pb.ListReq{
				Limit: params.Limit,
				Page:  params.Page,
			},
			ToTime: params.ListTime,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list todos", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
