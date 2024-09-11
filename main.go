package main

import (
    "log"
    "net/http"
    "job_task/routing"
    _ "github.com/lib/pq"
    "job_task/data" 
)

func main() {
    // Подключаемся к базе данных
    data.ConnectDB()

    // Настраиваем маршрутизацию запросов
    http.HandleFunc("/issue-tokens", routing.IssueTokens) // Обрабатываем запросы на создание токенов
    http.HandleFunc("/refresh-tokens", routing.RefreshTokens) // Обрабатываем запросы на обновление токенов

    // Запускаем сервер и слушаем на порту 8080
    log.Println("Server running on :8080") // Логируем сообщение о запуске сервера
    http.ListenAndServe(":8080", nil) // Запускаем сервер и ждем запросы
}
