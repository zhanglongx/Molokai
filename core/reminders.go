package core

import (
	"crypto/tls"
	"errors"
	"log"

	gomail "gopkg.in/mail.v2"

	"github.com/fxtlabs/date"
	"github.com/zhanglongx/Molokai/runner"
)

type Reminders struct {
	// Reminder to
	To string `yaml:"to"`

	messages []runner.Result `yaml:"-"`
}

var (
	errReminderCfg = errors.New("reminder cfg error")
)

func (r *Reminders) Init() error {
	if smtpCfg.From.User == "" {
		return errReminderCfg
	}

	return nil
}

func (r *Reminders) Fill(msg runner.Result) {
	r.messages = append(r.messages, msg)
}

func (r *Reminders) Send() error {
	var body string
	for _, m := range r.messages {
		body += m.Message + "\n"
	}

	if body == "" {
		return nil
	}

	cfg := &smtpCfg

	m := gomail.NewMessage()

	// Set E-Mail Header
	// TODO: using TZ/Shanghai ?
	m.SetHeader("From", cfg.From.User)
	m.SetHeader("To", r.To)
	m.SetHeader("Subject", date.TodayUTC().String())
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer(cfg.Smtp.Smtp, cfg.Smtp.Port, cfg.From.User, cfg.From.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	d.TLSConfig = &tls.Config{ServerName: cfg.Smtp.Smtp}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Printf("send to %s: %s", r.To, body)
	return nil
}
