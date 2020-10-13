package utils

import (
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

// SendEmail : sends a email to destinatary
func SendEmail(destinataryEmail string, otp string) {
	emailAddress := os.Getenv("EMAIL_ACCOUNNT")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	protocol := os.Getenv("EMAIL_PROTOCOL")
	protocolPort := os.Getenv("EMAIL_PROTOCOL_PORT")

	e := email.NewEmail()
	e.From = "Arreglapp <" + emailAddress + ">"
	e.To = []string{destinataryEmail}
	e.Subject = "Codigo"
	// e.Text = []byte("Activa tu cuenta")
	e.HTML = []byte("<p>Utiliza el siguente codigo para seguir operando en arreglapp:  " + otp + " </p>")
	e.Send(protocol+":"+protocolPort, smtp.PlainAuth("", emailAddress, emailPassword, protocol))
}
