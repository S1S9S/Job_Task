package data

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

// Claims - структура для хранения данных о пользователе в JWT
type Claims struct {
    UserID string `json:"user_id"`  
    AccID  string `json:"acc_id"`   
    IP     string `json:"ip"`       
    jwt.StandardClaims // Стандартные поля JWT, такие как срок действия и тому подобгые 
}

// RefreshToken - структура для хранения данных о рефреш-токене
type RefreshToken struct {
    Hash      string    
    AccID     string    
    UserID    string    
    IP        string    
    ExpiresAt time.Time 
}
