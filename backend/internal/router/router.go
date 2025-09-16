package router

import (
	"project/internal/ent"
	controllers "project/internal/handlers"
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
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(client)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	csvRepository := repositories.NewCsvRepository(client)
	filepath := "./data/sample_data_100000.csv"
	csvService := services.NewCsvService(csvRepository, filepath)
	csvController := controllers.NewCsvController(csvService)

	// ルーターの作成
	r := gin.Default()

	r.Use(cors.Default())

	// Swaggerの設定
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			eg.GET("/hello", controllers.HelloWorld)
		}

		// ルーティンググループの作成
		itemRouter := v1.Group("/items")
		itemRouterWithAuth := v1.Group("/items", middlewares.AuthMiddleware(authService))
		authRouter := v1.Group("/auth")
		csvRouter := v1.Group("/csv")

		// ルーティングの設定
		itemRouter.GET("", itemController.FindAll)

		itemRouterWithAuth.GET("/:id", itemController.FindById)
		itemRouterWithAuth.POST("", itemController.Create)
		itemRouterWithAuth.PUT("/:id", itemController.Update)
		itemRouterWithAuth.DELETE("/:id", itemController.Delete)

		authRouter.POST("/signup", authController.SignUp)
		authRouter.POST("/login", authController.Login)

		csvRouter.POST("/process", csvController.ProcessCsv)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", controllers.HelloWorld)

	return r
}
