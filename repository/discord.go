package repository

import (
	"bytes"
	"github.com/cockroachdb/errors"
	"github.com/usagiga/yagisan/model"
	"net/http"
)

type DiscordRepositoryImpl struct {
	client http.Client
}

func NewDiscordRepository(client http.Client) DiscordRepository {
	return &DiscordRepositoryImpl{
		client: client,
	}
}

func (repo *DiscordRepositoryImpl) SendTextChat(addr *model.Address, content string) (err error) {
	client := repo.client

	// Build request
	req, err := repo.buildRequest(addr, content)
	if err != nil {
		return errors.Wrap(err, "failed to build request")
	}

	// Request
	_, err = client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}

	// Handle Response
	// TODO: handle response

	return nil
}

func (repo *DiscordRepositoryImpl) buildRequest(addr *model.Address, content string) (req *http.Request, err error) {
	url := addr.WebhookURL.String()
	bodyBuf := bytes.NewBufferString(content)
	req, err = http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	return req, nil
}
