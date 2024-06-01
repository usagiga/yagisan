package model

type Mail struct {
	From string
	To   string // TODO: compatible with multiple Envelope To
	Body string
}

// String returns formatted mail
func (mail *Mail) String() (content string) {
	// TODO: Use optimal format
	return mail.Body
}
