package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/VanThen60hz/GoShop/global"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTemplateEmailOtp(
	to []string, from string, htmlTemplate string,
	dataTemplate map[string]interface{},
) error {
	htmlBody, err := getMailTemplate(htmlTemplate, dataTemplate)
	if err != nil {
		global.Logger.Error("Get mail template error: ", zap.Error(err))
		return err
	}

	return send(to, from, htmlBody)
}

func getMailTemplate(templateName string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t, err := template.New(templateName).ParseFiles("templates-email/" + templateName)
	if err != nil {
		return "", err
	}

	err = t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}

	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	smtpConfig := global.Config.SMTP

	contentEmail := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Test",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	messageMail := BuildMessage(contentEmail)

	authentication := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)
	err := smtp.SendMail(smtpConfig.Host+":"+smtpConfig.Port, authentication, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Send email error: ", zap.Error(err))
		return err
	}

	return nil
}
