package routes

import (
	"tugas-oti/controller"
	middlewares "tugas-oti/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	updatePassMiddlewareRoute := r.Group("update-password")
	updatePassMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	updatePassMiddlewareRoute.PUT("/:id", controller.UpdatePassword)

	r.GET("/event", controller.GetAllEvent)
	r.GET("/:id", controller.GetEventByID)
	eventMiddlewareRoute := r.Group("/event")
	eventMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	eventMiddlewareRoute.POST("/", controller.CreateEvent)
	eventMiddlewareRoute.PUT("/:id", controller.UpdateEvent)
	eventMiddlewareRoute.DELETE("/:id", controller.DeleteEvent)

	r.GET("/enrollment", controller.GetAllEnrollment)
	enrollmentMiddlewareRoute := r.Group("/enrollment")
	enrollmentMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	enrollmentMiddlewareRoute.POST("/", controller.CreateEnrollment)
	enrollmentMiddlewareRoute.DELETE("/:id", controller.DeleteEnrollment)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
