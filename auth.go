package twitter

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"twitter/pkg/errors"
)

var (
	UsernameMinLength = 2
	PasswordMinLength = 6
)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type RegisterInput struct {
	Email             string
	Password          string
	ConfirmedPassword string
	Username          string
}
type LoginInput struct {
	Email    string
	Password string
}
type AuthResponse struct {
	AccessToken  string
	RefreshToken string
	User         User
}
type AuthServiceInterface interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse, error)
	Login(ctx context.Context, input LoginInput) (AuthResponse, error)
}

func (in *RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
	in.Username = strings.TrimSpace(in.Username)
}

func (in *RegisterInput) Validate() error {
	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf("%w: username too short, (%d) characters at least", errors.ErrValidation, UsernameMinLength)
	}
	if !emailRegex.MatchString(in.Email) {
		return fmt.Errorf("%w: invalid email format", errors.ErrValidation)
	}
	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: password too short, (%d) characters at least", errors.ErrValidation, PasswordMinLength)
	}
	if in.Password != in.ConfirmedPassword {
		return fmt.Errorf("%w: passwords do not match", errors.ErrValidation)
	}
	return nil
}

func (in *LoginInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
}

func (in *LoginInput) Validate() error {
	if !emailRegex.MatchString(in.Email) {
		return fmt.Errorf("%w: invalid email format", errors.ErrValidation)
	}
	if len(in.Password) < 1 {
		return fmt.Errorf("%w:  Password Required", errors.ErrValidation)
	}
	return nil
}
