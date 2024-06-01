package module

import (
	"github.com/cockroachdb/errors"
	"github.com/usagiga/yagisan/model"
	"github.com/usagiga/yagisan/repository"
)

type MailModuleImpl struct {
	addrModule  AddressModule
	discordRepo repository.DiscordRepository
}

func NewMailModule(
	addrModule AddressModule,
	discordRepo repository.DiscordRepository,
) MailModule {
	return &MailModuleImpl{
		addrModule:  addrModule,
		discordRepo: discordRepo,
	}
}

func (mod *MailModuleImpl) Deliver(mail *model.Mail) (err error) {
	// infer address
	addr, err := mod.addrModule.InferForwardingAddress(mail)
	if err != nil {
		return errors.Wrap(err, "failed to infer forwarding address")
	}

	// send
	err = mod.discordRepo.SendTextChat(addr, mail.String())
	if err != nil {
		return errors.Wrap(err, "failed to deliver mail into discord")
	}

	return nil
}
