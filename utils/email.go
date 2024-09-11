package utils

import (
    "log"
    "net/smtp"
)

// SendEmailWarning отправляет уведомление по электронной почте о смене IP-адреса
func SendEmailWarning(userID string, newIP string) {
    from := "noreply@example.com" 
    pass := "password"           
    to := "user@example.com"    

    // Формируем сообщение электронной почты
    msg := "Subject: IP Change Alert\n\nYour IP changed to: " + newIP

    // Отправляем электронное письмо
    err := smtp.SendMail("smtp.example.com:587",         // Адрес и порт SMTP сервера
        smtp.PlainAuth("", from, pass, "smtp.example.com"),// Данные для аутентификации 
        from,             
        []string{to},     
        []byte(msg))      

    // Проверяем, произошла ли ошибка при отправке электронной почты
    if err != nil {
        log.Println("Ошибка отправки электронной почты:", err) // Логируем ошибку, если отправка письма не удалась
    }
}
