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
    subs := r.Group("/subcategories")
	{
		subs.PATCH("/:id", h.UpdateSub)
		subs.DELETE("/:id", h.DeleteSub)

	}
}

// RegisterNestedRoutes вешает маршруты под /categories
func (h *SubcategoryHandler) RegisterNestedRoutes(categories *gin.RouterGroup) {
	categories.POST("/:id/subcategories", h.CreateSub)
	categories.GET("/:id/subcategories", h.GetSuBByID)
}

func (h *SubcategoryHandler) CreateSub(c *gin.Context) {
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

func (h *SubcategoryHandler) GetSuBByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	subcat, err := h.service.GetSubcategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcat)
}

func (h *SubcategoryHandler) UpdateSub(c *gin.Context) {
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

func (h *SubcategoryHandler) DeleteSub(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteSubcategory(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "subcategory deleted"})
}

func (h *SubcategoryHandler) ListSubByCategory(c *gin.Context) {
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
