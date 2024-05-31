package smtp

import (
	"github.com/emersion/go-smtp"
)

type BackendImpl struct {
	authenticator Authenticator
	handler       Handler
}

func NewBackend(
	authenticator Authenticator,
	handler Handler,
) smtp.Backend {
	return &BackendImpl{
		authenticator: authenticator,
		handler:       handler,
	}
}

func (b *BackendImpl) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return NewAuthSession(b.authenticator, b.handler), nil
}
