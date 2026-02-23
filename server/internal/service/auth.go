package service

import (
	"errors"
	"time"

	"github.com/VinceZCL/FinalYearProject/app/config"
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model"
	"github.com/VinceZCL/FinalYearProject/types/model/dto"
	"github.com/VinceZCL/FinalYearProject/types/model/param"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var JWTSecretKey = []byte(config.Get().Security.Secretkey)

type Claims struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Type     string `json:"type"`
	Status   string `json:"status"`
	jwt.RegisteredClaims
}

type AuthService struct {
	authRepo repository.AuthRepository
	userRepo repository.UserRepository
}

func NewAuthService(authRepo repository.AuthRepository, userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(c echo.Context, req param.NewUser) (*dto.User, error) {
	var userType string
	if req.Type != "" {
		userType = req.Type
	} else {
		userType = "user"
	}
	hashed, err := tools.HashPass(req.Password)
	if err != nil {
		return nil, err
	}
	newReq := param.NewUser{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Type:     userType,
	}
	input := model.User{
		Name:     newReq.Name,
		Email:    newReq.Email,
		Password: newReq.Password,
		Type:     newReq.Type,
		Status:   "active",
	}

	user, err := s.userRepo.NewUser(input)
	if err != nil {
		c.Logger().Errorf("Service | AuthService | Register: %w", err)
		return nil, err
	}

	dto := &dto.User{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Type:   user.Type,
		Status: user.Status,
	}
	return dto, nil
}

func (s *AuthService) Login(c echo.Context, param param.Login) (string, error) {
	user, err := s.authRepo.GetCredentials(param.Email)
	if err != nil {
		c.Logger().Errorf("Service | AuthService | GetCredentials (%s): %w", param.Email, err)
		return "", errors.New("Email not found")
	}
	if !tools.ComparePass(user.Password, param.Password) {
		return "", errors.New("Incorrect Password")
	}
	token, err := tokenGen(c, user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func tokenGen(c echo.Context, user *model.User) (string, error) {
	if user.Status != "active" {
		return "", errors.New("User deactivated")
	}

	expire := time.Now().Add(24 * time.Hour)

	claims := Claims{
		UserID:   user.ID,
		Username: user.Name,
		Email:    user.Email,
		Type:     user.Type,
		Status:   user.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "github.com/VinceZCL/FinalYearProject",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		c.Logger().Errorf("Service | AuthService | tokenGen: %w", err)
		return "", err
	}
	return tokenString, nil
}
