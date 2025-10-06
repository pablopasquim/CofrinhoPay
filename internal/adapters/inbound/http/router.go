package http

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(r *gin.Engine) {
	// TODO: Implementar rotas quando os handlers estiverem prontos

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "CofrinhoPay API is running",
		})
	})

	// TODO: Adicionar grupos de rotas
	// v1 := r.Group("/api/v1")
	// {
	//     users := v1.Group("/users")
	//     {
	//         users.POST("/", userHandler.CreateUser)
	//         users.POST("/login", userHandler.LoginUser)
	//         users.GET("/:id", userHandler.GetUser)
	//         users.PUT("/:id", userHandler.UpdateUser)
	//         users.DELETE("/:id", userHandler.DeleteUser)
	//     }
	//
	//     categories := v1.Group("/categories")
	//     {
	//         categories.POST("/", categoryHandler.CreateCategory)
	//         categories.GET("/", categoryHandler.GetAllCategories)
	//         categories.GET("/:id", categoryHandler.GetCategory)
	//         categories.PUT("/:id", categoryHandler.UpdateCategory)
	//         categories.DELETE("/:id", categoryHandler.DeleteCategory)
	//     }
	//
	//     accounts := v1.Group("/accounts")
	//     {
	//         accounts.POST("/", accountHandler.CreateAccount)
	//         accounts.GET("/", accountHandler.GetAllAccounts)
	//         accounts.GET("/:id", accountHandler.GetAccount)
	//         accounts.PUT("/:id", accountHandler.UpdateAccount)
	//         accounts.DELETE("/:id", accountHandler.DeleteAccount)
	//     }
	//
	//     transactions := v1.Group("/transactions")
	//     {
	//         transactions.POST("/", transactionHandler.CreateTransaction)
	//         transactions.GET("/", transactionHandler.GetAllTransactions)
	//         transactions.GET("/:id", transactionHandler.GetTransaction)
	//         transactions.PUT("/:id", transactionHandler.UpdateTransaction)
	//         transactions.DELETE("/:id", transactionHandler.DeleteTransaction)
	//     }
	// }
}
