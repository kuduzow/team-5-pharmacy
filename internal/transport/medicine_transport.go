package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/repository"
	"github.com/kuduzow/team-5-pharmacy/internal/services"
)

type MedicineHendler struct {
	service services.MedicineService
}

func NewHandlerMedicine(services services.MedicineService) *MedicineHendler {
	return &MedicineHendler{service: services}
}
func (h *MedicineHendler) RegisterRoutes(r *gin.Engine) {
	medicine := r.Group("/medicine")
	{
		medicine.POST("", h.Create)
        medicine.GET("", h.List)
        medicine.GET("/:id", h.GetByID)
        medicine.PATCH("/:id", h.Update)
        medicine.DELETE("/:id", h.Delete)

	}

}
func (h *MedicineHendler) Create(c *gin.Context) {
	var req models.CreateMedicineRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	medicine, err := h.service.CreateMedicine(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, medicine)

}
func (h *MedicineHendler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	medicine, err := h.service.GetMedecinesById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHendler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req models.UpdateMedicineRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	medicine, err := h.service.UpdateMedicine(uint(id), &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHendler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.service.DeleteMedecineById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *MedicineHendler) List(c *gin.Context) {
	var filter repository.MedicinesFilter

	if s := c.Query("category_id"); s != "" {
		if v, err := strconv.ParseUint(s, 10, 64); err == nil {
			x := uint(v)
			filter.CategoryId = &x
		}
	}
	if s := c.Query("subcategory_id"); s != "" {
		if v, err := strconv.ParseUint(s, 10, 64); err == nil {
			x := uint(v)
			filter.SubcategoryId = &x
		}
	}
	if s := c.Query("in_stock"); s != "" {
		if b, err := strconv.ParseBool(s); err == nil {
			filter.InStock = &b
		}
	}

	medicines, err := h.service.List(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medicines)
}