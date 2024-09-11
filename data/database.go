package data

import (
    "database/sql"
    "log"
)

var DB *sql.DB

// ConnectDB подключается к базе данных PostgreSQL
func ConnectDB() {
    // Здесь задаем строку подключения к базе данных
    connStr := "user=zalupiy password=zalupiy228 dbname=Tron sslmode=disable"
    
    var err error
    // Открываем подключение к базе данных
    DB, err = sql.Open("postgres", connStr)
    
    // Если возникла ошибка при открытии подключения, выводим ее в лог
    if err != nil {
        log.Println("DB error") // Выводим сообщение об ошибке, если не удается подключиться
        return // Выходим из функции, так как подключение не удалось
    }
    
    // Подключение прошло успешно, дальше можно работать с базой данных
}
