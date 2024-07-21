package pkg

import (
        "log"

        "user_manager/config"
        //"github.com/gofiber/fiber"
        "gopkg.in/gomail.v2"
)

type EmailAdpater struct {
        To      string `json:"to"`
        Subject string `json:"subject"`
        Body    string `json:"body"`
}

func (data *EmailAdpater) SendMessageViaEmail() error {
        // Compose the email
        email := gomail.NewMessage()
        email.SetHeader("From", config.ReadConfigEnvVariable("Tele_email"))
        email.SetHeader("To", data.To)
        email.SetHeader("Subject", data.Subject)
        email.SetBody("text/plain", data.Body)
        log.Println(data.To)
        log.Println(config.ReadConfigEnvVariable("Tele_email"))
        // Send the email
        dialer := gomail.NewDialer(config.ReadConfigEnvVariable("HOST"), 465, config.ReadConfigEnvVariable("Tele_email"), config.ReadConfigEnvVariable("PASSWORD"))
        if err := dialer.DialAndSend(email); err != nil {
                log.Printf("Failed to send email: %v", err)
                return err
        }
        log.Println("Email message sent successfully!")
        return nil

}

