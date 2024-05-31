package smtp

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/google/uuid"
	"io"
	"log/slog"
)

const (
	ContextKey_RequestID string = "request-id"
	ContextKey_From             = "from"
	ContextKey_To               = "to"
)

type SessionImpl struct {
	ctx           context.Context
	authenticator Authenticator
	handler       Handler
}

func NewAuthSession(
	authenticator Authenticator,
	handler Handler,
) smtp.Session {
	session := &SessionImpl{
		authenticator: authenticator,
		handler:       handler,
	}
	session.initialize()

	return session
}

func (s *SessionImpl) initialize() {
	ctx := context.Background()

	// Set Request ID into context
	reqId := uuid.Must(uuid.NewV7())
	ctx = context.WithValue(ctx, ContextKey_RequestID, reqId)
	ctx = context.WithValue(ctx, ContextKey_From, "")
	ctx = context.WithValue(ctx, ContextKey_To, "")
	slog.DebugContext(ctx, fmt.Sprintf("request-id: %s", reqId))
	s.ctx = ctx

	// inject values into receiver
	s.ctx = ctx
}

func (s *SessionImpl) AuthMechanisms() []string {
	return s.authenticator.AuthMechanisms()
}

func (s *SessionImpl) Auth(mech string) (sasl.Server, error) {
	server, err := s.authenticator.Auth(mech)
	if err != nil {
		return nil, errors.Wrap(err, "authentication failed")
	}

	return server, nil
}

func (s *SessionImpl) Mail(from string, opts *smtp.MailOptions) error {
	// set Envelope From into context
	ctx := s.ctx
	ctx = context.WithValue(ctx, ContextKey_From, from)
	slog.DebugContext(ctx, fmt.Sprintf("from: %s", from))
	s.ctx = ctx

	return nil
}

func (s *SessionImpl) Rcpt(to string, opts *smtp.RcptOptions) error {
	// set Envelope(Rcpt) To into context
	ctx := s.ctx
	ctx = context.WithValue(ctx, ContextKey_To, to)
	slog.DebugContext(ctx, fmt.Sprintf("to: %s", to))
	s.ctx = ctx

	return nil
}

func (s *SessionImpl) Data(r io.Reader) error {
	// TODO: use goroutine to handle RSET, Logout properly

	// read all DATA
	content, err := io.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "failed to read DATA")
	}

	// read From / To
	// TODO: error handling
	ctx := s.ctx
	from, _ := ctx.Value(ContextKey_From).(string)
	to, _ := ctx.Value(ContextKey_To).(string)

	// call handler
	err = s.handler.HandleMail(ctx, from, string(content), to)
	if err != nil {
		return errors.Wrap(err, "error raised in handler")
	}

	return nil
}

func (s *SessionImpl) Reset() {
	slog.DebugContext(s.ctx, "session reset")
	s.initialize()
}

func (s *SessionImpl) Logout() error {
	slog.DebugContext(s.ctx, "session logout")

	// NOP
	return nil
}
