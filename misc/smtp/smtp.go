package smtp

import (
	"github.com/emersion/go-smtp"
)

// NewSmtpServer construct smtp server of go-smtp.
func NewSmtpServer() (server *smtp.Server) {
	// TODO: configure server here
	server = smtp.NewServer(nil)
	return server
}
