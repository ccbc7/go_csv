package router

import (
	"project/internal/ent"
	"project/internal/handlers"
	"project/internal/middlewares"
	"project/internal/repositories"
	"project/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"project/docs"
)

func SetupRouter(client *ent.Client) *gin.Engine {
	itemRepository := repositories.NewItemRepository(client)
	itemService := services.NewItemService(itemRepository)
	itemHandler := handlers.NewItemHandler(itemService)

	authRepository := repositories.NewAuthRepository(client)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	csvRepository := repositories.NewCsvRepository(client)
	filepath := "./data/sample_data_100000.csv"
	csvService := services.NewCsvService(csvRepository, filepath)
	csvHandler := handlers.NewCsvHandler(csvService)

	// ルーターの作成
	r := gin.Default()

	r.Use(cors.Default())

	// Swaggerの設定
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			eg.GET("/hello", handlers.HelloWorld)
		}

		// ルーティンググループの作成
		itemRouter := v1.Group("/items")
		itemRouterWithAuth := v1.Group("/items", middlewares.AuthMiddleware(authService))
		authRouter := v1.Group("/auth")
		csvRouter := v1.Group("/csv")

		// ルーティングの設定
		itemRouter.GET("", itemHandler.FindAll)

		itemRouterWithAuth.GET("/:id", itemHandler.FindById)
		itemRouterWithAuth.POST("", itemHandler.Create)
		itemRouterWithAuth.PUT("/:id", itemHandler.Update)
		itemRouterWithAuth.DELETE("/:id", itemHandler.Delete)

		authRouter.POST("/signup", authHandler.SignUp)
		authRouter.POST("/login", authHandler.Login)

		csvRouter.POST("/process", csvHandler.ProcessCsv)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", handlers.HelloWorld)

	return r
}
