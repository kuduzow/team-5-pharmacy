package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/services"
)

type CategoryHandler struct {
	service services.CategoryService
    subcats *SubcategoryHandler
}

func NewCategoryHandler(service services.CategoryService, subcats *SubcategoryHandler) *CategoryHandler {
	return &CategoryHandler{service: service, subcats: subcats}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req models.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.service.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	category, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, services.ErrCategorytNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req models.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.service.UpdateCategory(uint(id), &req)
	if err != nil {
		if errors.Is(err, services.ErrCategorytNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный id"})
		return
	}

	if err := h.service.DeleteCategory(uint(id)); err != nil {
		if errors.Is(err, services.ErrCategorytNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.Status(http.StatusOK)
}

func (h *CategoryHandler) RegisterRoutes(r *gin.Engine) {
	categories := r.Group("/categories")
	{
		categories.POST("", h.Create)
		categories.GET("/:id", h.GetByID)
		categories.PATCH("/:id", h.Update)
		categories.DELETE("/:id", h.Delete)
       
        categories.POST("/subcategories",h.subcats.CreateSub)
        categories.GET("/:id/subcategories", h.subcats.ListSubByCategory)
        categories.PATCH("/subcategories/:id",h.subcats.UpdateSub)
        categories.DELETE("/subcategories/:id",h.subcats.DeleteSub)

       
	}
}

