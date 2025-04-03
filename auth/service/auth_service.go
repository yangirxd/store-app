package service

import (
	"github.com/yangirxd/store-app/auth/domain"
	"github.com/yangirxd/store-app/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(email, password string) (*domain.User, error) {
	user, err := domain.NewUser(email, password)
	if err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token, err := domain.GenerateJWT(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
