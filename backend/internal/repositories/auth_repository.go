package repositories

import (
	"context"
	"errors"

	"project/internal/ent"
	"project/internal/ent/user"
)

type IAuthRepository interface {
	CreateUser(name, loginID, password string) (*ent.User, error)
	FindUser(loginID string) (*ent.User, error)
}

type AuthRepository struct {
	client *ent.Client
}

func NewAuthRepository(client *ent.Client) IAuthRepository {
	return &AuthRepository{client: client}
}

func (r *AuthRepository) CreateUser(name, loginID, password string) (*ent.User, error) {
	return r.client.User.Create().
		SetName(name).
		SetLoginID(loginID).
		SetPassword(password).
		Save(context.Background())
}

func (r *AuthRepository) FindUser(loginID string) (*ent.User, error) {
	user, err := r.client.User.Query().
		Where(user.LoginIDEQ(loginID)).
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
