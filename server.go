package main

import (
	"dumbmerch-go/config"
	"dumbmerch-go/handler"
	"dumbmerch-go/middleware"
	"dumbmerch-go/repository"
	"dumbmerch-go/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	productRepository repository.ProductRepository = repository.NewProductRepo(db)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	jwtService        service.JWTService           = service.NewJWTService()
	productService    service.ProductService       = service.NewProductService(productRepository)
	userService       service.UserService          = service.NewUserService(userRepository)
	authHandler       handler.AuthHandler          = handler.NewAuthHandler(authService, jwtService, userService)
	productHandler    handler.ProductHandler       = handler.NewProductHandler(productService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	productRoutes := server.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/", productHandler.All)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	server.Run("localhost:8000")
}
