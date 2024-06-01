package handler

import (
	"context"
	"github.com/usagiga/yagisan/misc/smtp"
	"github.com/usagiga/yagisan/model"
	"github.com/usagiga/yagisan/module"
)

type SmtpHandlerImpl struct {
	mailModule module.MailModule
}

func NewSmtpHandler(
	mailModule module.MailModule,
) smtp.Handler {
	return &SmtpHandlerImpl{
		mailModule: mailModule,
	}
}

func (s *SmtpHandlerImpl) HandleMail(ctx context.Context, mail *model.Mail) (err error) {
	// TODO: validation mail here

	// deliver mail into discord
	err = s.mailModule.Deliver(mail)
	if err != nil {
		// TODO: error handling here
	}

	return nil
}
