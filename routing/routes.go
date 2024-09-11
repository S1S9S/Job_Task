package routing

import (
    "net/http"
    "job_task/handlers"
)

// IssueTokens - функция для маршрутизации запросов на выдачу токенов
func IssueTokens(w http.ResponseWriter, r *http.Request) {
    // Передаем запрос обработчику для выдачи токенов
    handlers.IssueTokens(w, r)
}

// RefreshTokens - функция для маршрутизации запросов на обновление токенов
func RefreshTokens(w http.ResponseWriter, r *http.Request) {
    // Передаем запрос обработчику для обновления токенов
    handlers.RefreshTokens(w, r)
}
