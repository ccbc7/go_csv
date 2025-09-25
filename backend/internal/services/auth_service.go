package services

import (
	"fmt"
	"os"
	"time"

	"project/internal/ent"
	"project/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(name, loginID, password string) error
	Login(loginID string, password string) (*string, error)
	GetUserFromToken(token string) (*ent.User, error)
}

type authService struct {
	repository repositories.AuthRepository
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{repository: repository}
}

func (s *authService) SignUp(name, loginID, password string) error {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.repository.CreateUser(name, loginID, string(hashedPassword))
	return err
}

func (s *authService) Login(loginID string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(loginID)
	if err != nil {
		return nil, err
	}

	// パスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// トークンの作成
	token, err := CreateToken(uint(foundUser.ID), foundUser.LoginID)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateToken(userId uint, loginID string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// subはsubjectの略で、JWTの主題を表す
		// loginIDはユーザーのログインID
		// expはJWTの有効期限を表す
		"sub":      userId,
		"login_id": loginID,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	// 署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (s *authService) GetUserFromToken(tokenString string) (*ent.User, error) {
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

	var user *ent.User
	// トークンのクレームを取得し、トークンが有効か確認
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// トークンの有効期限を確認
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// ErrTokenExpiredはjwtパッケージで定義されているエラー
			return nil, jwt.ErrTokenExpired
		}

		// クレームからログインIDを取得し、ユーザーを検索
		user, err = s.repository.FindUser(claims["login_id"].(string))
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	// トークンが無効な場合のエラーハンドリング
	return nil, fmt.Errorf("invalid token")
}
