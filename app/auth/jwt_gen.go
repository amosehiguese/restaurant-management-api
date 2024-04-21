package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Token struct {
	Access		string
	Refresh		string
}

func GenNewTokens(id uuid.UUID, role int32) (*Token, error) {
	accessToken, err := generateNewAccessToken(id, role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Token{
		Access: accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id uuid.UUID, role int32) (string, error) {
	secret := os.Getenv(("JWT_SECRET_KEY"))

	minCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{} 

	claims["id"] = id
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["eat"] = time.Now().Add(time.Minute * time.Duration(minCount)).Unix()


	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken() (string, error) {
	hash := sha256.New()
	refresh := os.Getenv("JWT_REFRESH_KEY") + time.Now().String()

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}

	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())

	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}

func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}