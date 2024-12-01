package router

import (
	"log"
	"os"
	"personal-finance-app/internal/handlers"
	"personal-finance-app/internal/platform/middleware"
	"personal-finance-app/internal/repository"
	"personal-finance-app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}

	gin.SetMode(ginMode)
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	r.POST("/users", userHandler.Register)
	r.POST("/users/login", userHandler.Login)

	authMiddleware := middleware.AuthMiddleware()
	authGroup := r.Group("/", authMiddleware)

	// Категории
	authGroup.GET("/categories", categoryHandler.GetCategories)
	authGroup.POST("/categories", categoryHandler.CreateCategory)
	authGroup.PUT("/categories/:id", categoryHandler.UpdateCategory)
	authGroup.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	// Транзакции
	authGroup.POST("/transactions", transactionHandler.CreateTransaction)
	authGroup.GET("/transactions", transactionHandler.GetTransactions)
	authGroup.GET("/transactions/:id", transactionHandler.GetTransaction)
	authGroup.PUT("/transactions/:id", transactionHandler.UpdateTransaction)
	authGroup.DELETE("/transactions/:id", transactionHandler.DeleteTransaction)

	return r
}
