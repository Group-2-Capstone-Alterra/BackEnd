package middlewares

import (
	"PetPalApp/app/configs"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(configs.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

// Generate token jwt
func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() //Token expires after 12 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.JWT_SECRET))

}

func ExtractTokenUserId(e echo.Context) (int, string, error) {
	header := e.Request().Header.Get("Authorization")
	if header == "" {
		return 0, "", errors.New("authorization header is empty")
	}

	headerToken := strings.Split(header, " ")
	token := headerToken[len(headerToken)-1]
	tokenJWT, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.JWT_SECRET), nil
	})

	if err != nil {
		e.Error(err)
		return 0, "", err
	}

	if tokenJWT.Valid {
		claims := tokenJWT.Claims.(jwt.MapClaims)
		userId, isValidUserId := claims["userId"].(float64)
		role, isValidRole := claims["role"].(string)
		if !isValidUserId || !isValidRole {
			e.Error(errors.New("jwt not found"))
			return 0, "", errors.New("jwt not found")
		}
		return int(userId), role, nil
	}

	e.Error(errors.New("jwt not found"))
	return 0, "", errors.New("jwt not found")
}
