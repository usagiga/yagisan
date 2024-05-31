package smtp

import "context"

type Handler interface {
	HandleMail(ctx context.Context, from string, content string, to string) (err error)
}
