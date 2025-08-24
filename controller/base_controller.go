package controller

import (
	"go-initial-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BaseController[T any] struct {
	service service.BaseServiceInterface[T]
}

func NewBaseController[T any](service service.BaseServiceInterface[T]) *BaseController[T] {
	return &BaseController[T]{service: service}
}

func (c *BaseController[T]) GetAll(ctx *gin.Context) {
	items, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *BaseController[T]) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	item, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *BaseController[T]) Create(ctx *gin.Context) {
	var item T
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := c.service.Create(item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *BaseController[T]) Update(ctx *gin.Context) {
	var item T
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := c.service.Update(item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *BaseController[T]) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var item T
	if err := c.service.Delete(uint(id), item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *BaseController[T]) HardDelete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var item T
	if err := c.service.HardDelete(uint(id), item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *BaseController[T]) Paginate(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	items, total, err := c.service.Paginate(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total": total, "data": items})
}

func (c *BaseController[T]) Search(ctx *gin.Context) {
	field := ctx.Query("field")
	keyword := ctx.Query("keyword")
	items, err := c.service.Search(field, keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *BaseController[T]) FindWithTrashed(ctx *gin.Context) {
	items, err := c.service.FindWithTrashed()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *BaseController[T]) OnlyTrashed(ctx *gin.Context) {
	items, err := c.service.OnlyTrashed()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *BaseController[T]) Restore(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var item T
	if err := c.service.Restore(uint(id), item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "restored"})
}
