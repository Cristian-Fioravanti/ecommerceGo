package routes

import(
	"github.com/cristian-fioravanti/ecommerceGo/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/cristian-fioravanti/ecommerceGo/middleware"

)
func UserRoutes(IncomingRoutes *gin.Engine, db *gorm.DB) {
	IncomingRoutes.POST("/users/signup", controllers.Signup(db))
	IncomingRoutes.POST("/users/login", controllers.Login(db))
	// IncomingRoutes.POST("/users/refresh", controllers.RefreshToken(db))

	IncomingRoutes.GET("/users/products/:productId", controllers.GetProductById(db))
	IncomingRoutes.GET("/users/products", controllers.GetProducts(db))
	
	adminRoutes := IncomingRoutes.Group("/admin")
	adminRoutes.Use(middleware.AdminMiddleware(db)) 

	adminRoutes.POST("/products", controllers.CreateProduct(db))
	adminRoutes.PUT("/products/:productId", controllers.UpdateProduct(db))
	adminRoutes.DELETE("/products/:productId", controllers.DeleteProduct(db))
}