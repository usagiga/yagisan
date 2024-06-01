package module

import "github.com/usagiga/yagisan/model"

type MailModule interface {
	Deliver(mail *model.Mail) (err error)
}

type AddressModule interface {
	RegisterAddresses(addrs ...*model.Address) (err error)
	InferForwardingAddress(mail *model.Mail) (forwarded *model.Address, err error)
}
