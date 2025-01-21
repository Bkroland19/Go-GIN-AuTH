package utils

import (
	"errors"

	"strconv"

	"time"

	"github.com/golang-jwt/jwt/v5"
)


const secretKey = "yahamachatecxwre"

func GenerateToken(email string ,userId int64) (string , error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(), 
	})


	return token.SignedString([]byte(secretKey))
}



func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("failed to parse token")
	}

	// Validate the token
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	// Extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Extract userId from claims
	userIdRaw, ok := claims["userId"]
	if !ok {
		return 0, errors.New("userId not found in token claims")
	}

	// Handle userId as either float64 or string
	var userId int64
	switch v := userIdRaw.(type) {
	case float64:
		userId = int64(v)
	case string:
		parsedId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, errors.New("userId in token claims is not a valid number")
		}
		userId = parsedId
	default:
		return 0, errors.New("unexpected userId type in token claims")
	}

	return userId, nil
}

