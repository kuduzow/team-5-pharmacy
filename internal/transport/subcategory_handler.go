package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/services"
)

type SubcategoryHandler struct {
	service services.SubcategoryService
}

func NewSubcategoryHandler(s services.SubcategoryService) *SubcategoryHandler {
	return &SubcategoryHandler{service: s}
}

func (h *SubcategoryHandler) RegisterRoutes(r *gin.Engine) {
	subcats := r.Group("/subcategories")
	{
		subcats.POST("", h.Create)
		subcats.GET("/:id", h.GetByID)
		subcats.PUT("/:id", h.Update)
		subcats.DELETE("/:id", h.Delete)
		subcats.GET("/category/:category_id", h.ListByCategory)
	}
}

func (h *SubcategoryHandler) Create(c *gin.Context) {
	var req models.CreateSubcategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	subcat, err := h.service.CreateSubcategory(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcat)
}

func (h *SubcategoryHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	subcat, err := h.service.GetSubcategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcat)
}

func (h *SubcategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.UpdateSubcategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	subcat, err := h.service.UpdateSubcategory(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcat)
}

func (h *SubcategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteSubcategory(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "subcategory deleted"})
}

func (h *SubcategoryHandler) ListByCategory(c *gin.Context) {
	cabIDStr := c.Param("category_id")

	id, err := strconv.ParseUint(cabIDStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subcats, err := h.service.ListSubcategoriesByCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcats)
}
