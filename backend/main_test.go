package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"project/dto"
	"project/infra"
	"project/models"
	"project/services"
)

func TestMain(m *testing.M) {
	// テスト用の.envファイルを読み込む
	if err := godotenv.Load(".env.test"); err != nil {
		log.Println("Error loading .env.test file:", err)
	}

	// テスト用のデータベースをセットアップ
	code := m.Run()

	// テストが終わったらテスト用のデータベースを削除
	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム2", Price: 1000, Description: "テスト２", SoldOut: false, UserID: 1},
		{Name: "テストアイテム3", Price: 1000, Description: "テスト３", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Item{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

// t *testing.T はテストの状態と結果を報告するためのオブジェクト
func TestFindAll(t *testing.T) {
	// テスト用のデータをセットアップ
	router := setup()

	// HTTPレスポンスを記録するオブジェクトを作成
	w := httptest.NewRecorder()

	// NewRequestを使ってリクエストを作成
	req, _ := http.NewRequest("GET", "/items", nil)

	// ServeHTTPメソッドを使ってリクエストを実行
	router.ServeHTTP(w, req)

	// resの型を定義しているだけで、中身は空のmap
	var res map[string][]models.Item

	// レスポンスのボディをJSON形式からresで定義した構造体に格納
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーションを使ってテストの結果を確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスのボディに含まれるデータの数が3であることを確認
	assert.Equal(t, 3, len(res["data"]))
	fmt.Println(res)
}

func TestCreate(t *testing.T) {
	router := setup()

	// サービス層のCreateTokenメソッドを使ってトークンを作成
	token, err := services.CreateToken(1, "test1@example.com")

	// tはテストの状態と結果の報告用オブジェクト, nilはエラーがないことを示す, errは実際のエラー
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム4",
		Price:       1000,
		Description: "Createテスト",
	}

	// createItemInputをJSON形式に変換
	reqBody, _ := json.Marshal(createItemInput)

	// レコーダーを作成
	w := httptest.NewRecorder()

	// リクエストを作成
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))

	// リクエストヘッダーにトークンをセット
	req.Header.Set("Authorization", "Bearer "+*token)

	fmt.Println(req.Header)

	router.ServeHTTP(w, req)

	// レスポンスのボディを格納する変数を定義
	var res map[string]models.Item

	// レスポンスのボディをJSON形式からresで定義した構造体に格納
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーションを使ってテストの結果を確認
	assert.Equal(t, http.StatusCreated, w.Code)

	assert.Equal(t, uint(4), res["data"].ID)
}

func TestCreateUnauthorized(t *testing.T) {
	router := setup()

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム4",
		Price:       1000,
		Description: "Createテスト",
	}

	// createItemInputをJSON形式に変換
	reqBody, _ := json.Marshal(createItemInput)

	// レコーダーを作成
	w := httptest.NewRecorder()

	// リクエストを作成
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))

	fmt.Println(req.Header)

	router.ServeHTTP(w, req)

	// レスポンスのボディを格納する変数を定義
	var res map[string]models.Item

	// レスポンスのボディをJSON形式からresで定義した構造体に格納
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーションを使ってテストの結果を確認
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
