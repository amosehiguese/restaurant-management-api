package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateJWT(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func ValidateAdminRoleJWT(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}

	return errors.New("invalid admin token provided")
}

func ValidateUserRoleJWT(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 1 {
		return nil
	}

	return errors.New("invalid author token provided")
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (any, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil 
}

type TokenMetadata struct {
	UserID			uuid.UUID
	EAT				int64
}

func ExtractTokenMetadata(r *http.Request) (*TokenMetadata, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		eat := int64(claims["eat"].(float64))

		return &TokenMetadata{
			UserID: userID,
			EAT: eat ,
		}, nil
	}

	return nil, err
	
}