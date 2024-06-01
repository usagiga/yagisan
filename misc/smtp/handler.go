package smtp

import (
	"context"
	"github.com/usagiga/yagisan/model"
)

type Handler interface {
	HandleMail(ctx context.Context, mail *model.Mail) (err error)
}
