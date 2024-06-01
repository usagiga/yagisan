package handler

import (
	"context"
	"github.com/usagiga/yagisan/misc/smtp"
	"github.com/usagiga/yagisan/model"
)

type SmtpHandlerImpl struct{}

func NewSmtpHandler() smtp.Handler {
	return &SmtpHandlerImpl{}
}

func (s *SmtpHandlerImpl) HandleMail(ctx context.Context, mail *model.Mail) (err error) {
	//TODO implement me
	panic("implement me")
}
