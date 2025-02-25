package main

import (
	"flag"
	"fmt"
	"project/controllers"
	"project/infra"
	"project/middlewares"

	// "project/models"
	"project/repositories"
	"project/services"

	"project/database/seeders"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"project/docs"
	// "github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(db *gorm.DB) *gin.Engine {

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	csvRepository := repositories.NewCsvRepository(db)
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
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// CORSの設定

	r.GET("/", controllers.HelloWorld)

	// ルーティンググループの作成
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")
	csvRouter := r.Group("/csv")

	// ルーティングの設定
	itemRouter.GET("", itemController.FindAll)

	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.SignUp)
	authRouter.POST("/login", authController.Login)

	csvRouter.POST("/process", csvController.ProcessCsv)

	return r
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for a pet store.
// @BasePath /api/v1

func main() {
	seed := flag.Bool("seed", false, "Run the database seeders")
  // migrate := flag.Bool("migrate", false, "Run the database migrations")
	flag.Parse()
	// 初期化(環境変数の読み込み)
	infra.Initialize()

	// DB接続
	db := infra.SetupDB()

	if *seed {
		seeders.SeedAll(db)
		fmt.Println("Seeding completed")
	} else {
		fmt.Println("No operation specified")
	}

	// ルーターの設定 引数にDBを渡すことで、各レイヤー(サービス,リポジトリ,コントローラ)でDBを利用できる
	r := setupRouter(db)

	r.Run(":8080")
}

// itemRepository := repositories.NewItemMemoryRepository(items)
// items := []models.Item{
// 	{ID: 1, Name: "item1", Price: 100, Description: "This is item1", SoldOut: false},
// 	{ID: 2, Name: "item2", Price: 200, Description: "This is item2", SoldOut: false},
// 	{ID: 3, Name: "item3", Price: 300, Description: "This is item3", SoldOut: false},
// }
