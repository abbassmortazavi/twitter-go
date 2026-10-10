package domain

import (
	"context"
	"errors"
	"fmt"
	"twitter"
	errors2 "twitter/pkg/errors"

	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo twitter.UserRepo
}

func NewAuthService(userRepo twitter.UserRepo) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (as AuthService) Register(ctx context.Context, input twitter.RegisterInput) (twitter.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return twitter.AuthResponse{}, err
	}

	if _, err := as.UserRepo.GetByUsername(ctx, input.Username); !errors.Is(err, errors2.ErrNotFound) {
		return twitter.AuthResponse{}, twitter.ErrUsernameTaken
	}

	if _, err := as.UserRepo.GetByEmail(ctx, input.Email); !errors.Is(err, errors2.ErrNotFound) {
		return twitter.AuthResponse{}, twitter.ErrEmailTaken
	}

	user := twitter.User{
		Username: input.Username,
		Email:    input.Email,
	}
	//hashed
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return twitter.AuthResponse{}, fmt.Errorf("hashing password: %w", err)
	}
	user.Password = string(hashedPassword)

	user, err = as.UserRepo.Create(ctx, user)
	if err != nil {
		return twitter.AuthResponse{}, fmt.Errorf("creating user: %w", err)
	}
	return twitter.AuthResponse{
		AccessToken: "a token",
		User:        user,
	}, nil
}

func (as AuthService) Login(ctx context.Context, input twitter.LoginInput) (twitter.AuthResponse, error) {
	input.Sanitize()
	if err := input.Validate(); err != nil {
		return twitter.AuthResponse{}, err
	}
	user, err := as.UserRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, errors2.ErrNotFound):
			return twitter.AuthResponse{}, errors2.ErrBadCredentials
		default:
			return twitter.AuthResponse{}, fmt.Errorf("getting user: %w", err)
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return twitter.AuthResponse{}, errors2.ErrBadCredentials
	}
	data := twitter.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return twitter.AuthResponse{
		AccessToken: "a token",
		User:        data,
	}, err
}

//13
