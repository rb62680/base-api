package server

import (
	"github.com/dernise/base-api/controllers"
	"github.com/dernise/base-api/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"time"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "You successfully reached the base API."})
}

func (a API) SetupRouter() {
	router := a.Router

	router.Use(middlewares.ErrorMiddleware())
	router.Use(middlewares.CorsMiddleware(middlewares.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 := router.Group("/v1")
	{
		v1.GET("/", Index)
		users := v1.Group("/users")
		{
			userController := controllers.NewUserController(a.Database, a.EmailSender, a.Config)
			users.GET("/", userController.GetUsers)
			users.POST("/requestReset", userController.ResetPasswordRequest)
		}

		user := v1.Group("/user")
		{
			userController := controllers.NewUserController(a.Database, a.EmailSender, a.Config)
			user.GET("/:id", userController.GetUser)
			user.POST("/", userController.CreateUser)
			user.GET("/:id/activate/:activationKey", userController.ActivateUser)
			//users.GET("/:id/reset/:resetKey", userController.FormResetPassword)
			user.POST("/:id/reset/", userController.ResetPassword)
		}

		authentication := v1.Group("/auth")
		{
			authController := controllers.NewAuthController(a.Database, a.Config)
			authentication.POST("/", authController.Authentication)
		}

		authorized := v1.Group("/authorized")
		{
			authorized.Use(middlewares.AuthMiddleware())
			billing := authorized.Group("/billing")
			{
				billingController := controllers.NewBillingController(a.Database, a.EmailSender, a.Config)
				billing.POST("/", billingController.CreateTransaction)
			}
		}
	}
}
