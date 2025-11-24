package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/services"
)

type PaymentHandler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) RegisterRoutes(r *gin.Engine) {
	payments := r.Group("/payments")
	{
		payments.GET("", h.Create)
		payments.PATCH("", h.Update)
		payments.DELETE("/:id", h.Delete)
	}
}

func (h *PaymentHandler) Create(c *gin.Context) {
	var req models.CreatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.service.CreatePayment(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payment)

}

func (h *PaymentHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req models.UpdatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	payment, err := h.service.UpdatePayment(uint(id), &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.service.DeletePayment(uint(id)); err != nil {
		if errors.Is(err, services.ErrPaymentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	}
	c.Status(http.StatusOK)
}
