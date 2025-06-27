package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, role, secret string) (string, error) {
	// Создаем claims (данные в токене)
	claims := jwt.MapClaims{
		"sub":  userID,                                // ID пользователя
		"role": role,                                  // Роль пользователя
		"exp":  time.Now().Add(time.Hour * 72).Unix(), // Срок действия токена
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	return token.SignedString([]byte(secret))
}
