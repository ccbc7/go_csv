package services

import (
	"fmt"
	"os"
	"time"

	"gin-fleamarket/models"
	"gin-fleamarket/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error
	Login(email string, password string) (*string, error)
	GetUserFromToken(token string) (*models.User, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) SignUp(email string, password string) error {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	// パスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// トークンの作成
	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// subはsubjectの略で、JWTの主題を表す
		// emailはユーザーのメールアドレス
		// expはJWTの有効期限を表す
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	// 署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (s *AuthService) GetUserFromToken(tokenString string) (*models.User, error) {
	// トークンをパースして検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 型アサーションで署名方法を確認, 型があっていない場合はエラーを返す(HS256以外の署名方法は受け付けない)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 環境変数から秘密鍵を取得
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	var user *models.User
	// トークンのクレームを取得し、トークンが有効か確認
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// トークンの有効期限を確認
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// ErrTokenExpiredはjwtパッケージで定義されているエラー
			return nil, jwt.ErrTokenExpired
		}

		// クレームからメールアドレスを取得し、ユーザーを検索
		user, err = s.repository.FindUser(claims["email"].(string))
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	// トークンが無効な場合のエラーハンドリング
	return nil, fmt.Errorf("invalid token")
}
