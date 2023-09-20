package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/christhianjesus/priverion-challenge/internal/domain/user"
)

func NewUserService(repo user.UserRepository) user.UserService {
	return &userService{repo}
}

type userService struct {
	repo user.UserRepository
}

func (s *userService) CreateUser(ctx context.Context, data user.User) error {
	data.SetPassword(encryptSHA256(data.Password()))

	return s.repo.Create(ctx, data)
}

func (s *userService) AuthUser(ctx context.Context, username, password string) (user.User, error) {
	data, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if data.Password() != encryptSHA256(password) {
		return nil, errors.New("Invalid password")
	}

	return data, nil
}

func (s *userService) GetUser(ctx context.Context, userID string) (user.User, error) {
	return s.repo.Get(ctx, userID)
}

func (s *userService) UpdateUser(ctx context.Context, data user.User) error {
	data.SetPassword(encryptSHA256(data.Password()))

	return s.repo.Update(ctx, data)
}

func (s *userService) DeleteUser(ctx context.Context, userID string) error {
	return s.repo.Delete(ctx, userID)
}

func encryptSHA256(data string) string {
	bytesData := []byte(data)
	hash := sha256.Sum256(bytesData)

	return hex.EncodeToString(hash[:])
}
