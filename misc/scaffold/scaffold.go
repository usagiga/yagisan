package scaffold

import (
	"github.com/emersion/go-smtp"
	"github.com/usagiga/yagisan/handler"
	"github.com/usagiga/yagisan/middleware/authenticator"
	internalSmtp "github.com/usagiga/yagisan/misc/smtp"
)

// ScaffoldModules wires modules up using config
func ScaffoldModules() (server *smtp.Server) {
	auth := authenticator.NewTLSAuthenticator()
	h := handler.NewSmtpHandler()
	server = internalSmtp.NewSmtpServer(nil, auth, h) // TODO: inject config
	return server
}
