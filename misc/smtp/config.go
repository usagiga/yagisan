package smtp

import (
	"fmt"
	"github.com/emersion/go-smtp"
	"time"
)

type ServerConfig struct {
	Host              string
	Port              int
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	MaxMessageBytes   int64
	MaxRecipients     int
	AllowInsecureAuth bool

	EnableSMTPUTF8   bool
	EnableBINARYMIME bool
	EnableDSN        bool

	TLSConfig *TLSConfig
}

type TLSConfig struct {
	CertificateFilePath string
}

func NewServerConfig() *ServerConfig {
	// define default values
	return &ServerConfig{
		Host:            "localhost",
		Port:            10587,
		WriteTimeout:    10 * time.Second,
		ReadTimeout:     10 * time.Second,
		MaxMessageBytes: 1024 * 1024,
		MaxRecipients:   50,
	}
}

func (config *ServerConfig) Apply(server *smtp.Server) {
	server.Addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	server.Domain = config.Host
	server.WriteTimeout = config.WriteTimeout
	server.ReadTimeout = config.ReadTimeout
	server.MaxMessageBytes = config.MaxMessageBytes
	server.MaxLineLength = 2000 // from default value of library
	server.MaxRecipients = config.MaxRecipients
	server.AllowInsecureAuth = config.AllowInsecureAuth

	server.EnableSMTPUTF8 = config.EnableSMTPUTF8
	server.EnableREQUIRETLS = config.TLSConfig != nil
	server.EnableBINARYMIME = config.EnableBINARYMIME
	server.EnableDSN = config.EnableDSN
}
