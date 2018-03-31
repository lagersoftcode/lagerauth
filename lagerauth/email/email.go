package email

import (
	"bytes"
	"fmt"
	"html/template"
	"lagerauth/config"
	"net/smtp"

	"github.com/satori/go.uuid"
)

type EmailSender struct {
	Email string
	Pass  string
	Host  string
	Port  string
}

func New(c *config.Config) *EmailSender {
	return &EmailSender{c.EmailConfig.Email, c.EmailConfig.Pass, c.EmailConfig.Host, c.EmailConfig.Port}
}

func (es *EmailSender) Server() string {
	return es.Host + ":" + es.Port
}

func (es *EmailSender) SendResetCode(to string, code uuid.UUID) error {

	auth := smtp.PlainAuth("", es.Email, es.Pass, es.Host)
	template := &smtpTemplateData{
		From:    es.Email,
		To:      to,
		Subject: "Your lagersoft account password reset",
		Body:    fmt.Sprintf("Navigate to https://oauth.lagersoft.com/resetpass?code=%s to reset your password", code.String()),
	}

	return smtp.SendMail(es.Server(), auth, es.Email, []string{to}, template.Compile())
}

func (templateData *smtpTemplateData) Compile() []byte {
	t, _ := template.New("resetPassEmailTemplate").Parse(emailTemplate)

	var compiledTemplate bytes.Buffer
	t.Execute(&compiledTemplate, templateData)
	return compiledTemplate.Bytes()
}

type smtpTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}`
