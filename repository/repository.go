package repository

import "github.com/usagiga/yagisan/model"

type DiscordRepository interface {
	SendTextChat(addr *model.Address, content string) (err error)
}
