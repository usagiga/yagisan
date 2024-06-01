package module

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/usagiga/yagisan/model"
	"log/slog"
)

type AddressModuleImpl struct {
	addrMap map[string]*model.Address
}

func NewAddressModule() AddressModule {
	return &AddressModuleImpl{
		addrMap: make(map[string]*model.Address),
	}
}

func (mod *AddressModuleImpl) RegisterAddresses(addrs ...*model.Address) (err error) {
	beforeNumAddrs := len(mod.addrMap)

	for _, addr := range addrs {
		mod.addrMap[addr.ID] = addr
	}

	numAdded := len(mod.addrMap) - beforeNumAddrs
	// TODO: use context
	slog.Info(fmt.Sprintf("%d addresses registered"), numAdded)

	return nil
}

func (mod *AddressModuleImpl) InferForwardingAddress(mail *model.Mail) (forwarded *model.Address, err error) {
	// TODO: use username of to
	// in the case of `johndoe@example.com`, use "johndoe"
	to := mail.To
	found := mod.addrMap[to]
	if found == nil {
		return nil, errors.Errorf("address not found: %s", to)
	}

	return found, nil
}
