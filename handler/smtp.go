package handler

import (
	"context"
	"github.com/usagiga/yagisan/misc/smtp"
)

type SmtpHandlerImpl struct{}

func NewSmtpHandler() smtp.Handler {
	return &SmtpHandlerImpl{}
}

func (s *SmtpHandlerImpl) HandleMail(ctx context.Context, from string, content string, to string) (err error) {
	//TODO implement me
	panic("implement me")
}
