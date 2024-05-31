package smtp

import "github.com/emersion/go-sasl"

type Authenticator interface {
	AuthMechanisms() []string
	Auth(mech string) (sasl.Server, error)
}
