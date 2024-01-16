package services

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

// ! SendMail function send the email to the particular user
func SendMail(email string) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	mail_password := os.Getenv("MAIL_PASSWORD")

	auth := smtp.PlainAuth("", "yatinjal123@gmail.com", mail_password, "smtp.gmail.com")

	msg := "Your monthly expense reached 90% of your monthly limit"
	emails := []string{email}

	err=smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"yatinjal123@gmail.com",
		emails,
		[]byte(msg),
	)

	if err!=nil{
		log.Println("Error in sending the mail ",err)
		return err
	}
	return nil
}
