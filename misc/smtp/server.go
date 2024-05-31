package smtp

import (
	"github.com/emersion/go-smtp"
)

// NewSmtpServer construct smtp server of go-smtp.
func NewSmtpServer(
	config *ServerConfig,
	authenticator Authenticator,
	handler Handler,
) (server *smtp.Server) {
	backend := NewBackend(authenticator, handler)
	server = smtp.NewServer(backend)
	config.Apply(server)
	return server
}
