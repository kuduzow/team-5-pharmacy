package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/services"
)

func RegisterRoutes(router *gin.Engine,
	categoryService services.CategoryService,
	subcategoryService services.SubcategoryService,
	paymentService services.PaymentService,
	reviewService services.ReviewService,
	userService services.UserService,
) {
	// Инициализация хэндлеров
	categoryHandler := NewCategoryHandler(categoryService)
	subcategoryHandler := NewSubcategoryHandler(subcategoryService)
	paymentHandler := NewPaymentHandler(paymentService)
	reviewHandler := NewReviewHandler(reviewService)
	userHandler := NewUserHandler(userService)

	// Регистрация роутов
	categoryHandler.RegisterRoutes(router)
	subcategoryHandler.RegisterRoutes(router)
	paymentHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)

	// вложенные маршруты: /categories/:id/subcategories
	categoriesGroup := router.Group("/categories")
	subcategoryHandler.RegisterNestedRoutes(categoriesGroup)
}
