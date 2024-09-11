package tokens

import (
    "github.com/dgrijalva/jwt-go"
    "job_task/data"
    "errors"
)

var secret = []byte("secret") // Секретный ключ для подписи токенов

// CreateAccessToken создает новый JWT токен с указанными данными
func CreateAccessToken(claims *data.Claims) (string, error) {
    // Создаем новый токен с указанными данныи
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
    
    // Подписываем токен с использованием секрета
    tokenString, err := token.SignedString(secret)
    if err != nil {
        // Если произошла ошибка при подписании токена, возвращаем ошибку
        return "", err
    }
    
    // Возвращаем строку с токеном и nil как ошибку
    return tokenString, nil
}

// ValidateAccessToken проверяет валидность JWT токена
func ValidateAccessToken(tokenStr string) (*data.Claims, error) {
    // Парсим токен и извлекаем данные
    token, err := jwt.ParseWithClaims(tokenStr, &data.Claims{}, func(token *jwt.Token) (interface{}, error) {
        // Возвращаем секретный ключ для проверки подписи
        return secret, nil
    })
    if err != nil {
        // Если произошла ошибка при парсинге токена, возвращаем ошибку
        return nil, err
    }

    // Проверяем, что токен валиден и claims правильные
    claims, ok := token.Claims.(*data.Claims)
    if !ok || !token.Valid {
        // Если токен некорректен, возвращаем ошибку
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
