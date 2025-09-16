package main

import (
	"bytes"
	"context"
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

	"project/internal/dto"
	"project/internal/ent"
	"project/internal/router"
	"project/internal/services"
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

func setupTestData(client *ent.Client) {
	// テスト用のユーザーを作成
	user1, _ := client.User.Create().
		SetName("テストユーザー1").
		SetLoginID("test1@example.com").
		SetPassword("test1pass").
		Save(context.Background())

	user2, _ := client.User.Create().
		SetName("テストユーザー2").
		SetLoginID("test2@example.com").
		SetPassword("test2pass").
		Save(context.Background())

	// テスト用のアイテムを作成
	client.Item.Create().
		SetName("テストアイテム1").
		SetPrice(1000).
		SetDescription("").
		SetSoldOut(false).
		SetUserID(user1.ID).
		Save(context.Background())

	client.Item.Create().
		SetName("テストアイテム2").
		SetPrice(1000).
		SetDescription("テスト２").
		SetSoldOut(false).
		SetUserID(user1.ID).
		Save(context.Background())

	client.Item.Create().
		SetName("テストアイテム3").
		SetPrice(1000).
		SetDescription("テスト３").
		SetSoldOut(false).
		SetUserID(user2.ID).
		Save(context.Background())
}

func setup() *gin.Engine {
	// Entクライアントを作成
	dsn := "host=localhost user=postgres password=password dbname=test_db port=5432 sslmode=disable"
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// スキーマを作成
	client.Schema.Create(context.Background())

	setupTestData(client)
	return router.SetupRouter(client)
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
	var res map[string][]ent.Item

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
	var res map[string]ent.Item

	// レスポンスのボディをJSON形式からresで定義した構造体に格納
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーションを使ってテストの結果を確認
	assert.Equal(t, http.StatusCreated, w.Code)

	assert.Equal(t, 4, res["data"].ID)
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
	var res map[string]ent.Item

	// レスポンスのボディをJSON形式からresで定義した構造体に格納
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーションを使ってテストの結果を確認
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
