package auth

import (
	"net/smtp"

	"github.com/ajayjadhav201/common"
)

type MailService struct {
	mailAddr string
	from     string
	auth     smtp.Auth
}

func NewMailService(Username string, Password string) *MailService {
	// auth := smtp.PlainAuth("", Username, Password, "smtp.gmail.com")
	return &MailService{
		"",
		"",
		nil,
	}
}

// type Mail struct {
// 	From    string
// 	To      []string
// 	Subject string
// 	Body    string
// }

func (m *MailService) SendMail() error {
	if m == nil {
		//
		return common.Error("Internal server error")
	}
	//

	//
	to := []string{"vithaiautoindustries@gmail.com"}

	sub := "Subject: Reset password link\nClick on the belwo link to update your password"
	// body := "click on the below link to reset password"

	return smtp.SendMail("smtp.gmail.com:587", m.auth, m.from, to, []byte(sub))
}
