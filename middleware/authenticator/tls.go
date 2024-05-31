package authenticator

import (
	"github.com/emersion/go-sasl"
	"github.com/usagiga/yagisan/misc/smtp"
)

type TLSAuthenticator struct{}

func NewTLSAuthenticator() smtp.Authenticator {
	return &TLSAuthenticator{}
}

func (auth *TLSAuthenticator) AuthMechanisms() []string {
	//TODO implement me
	panic("implement me")
}

func (auth *TLSAuthenticator) Auth(mech string) (sasl.Server, error) {
	//TODO implement me
	panic("implement me")
}
