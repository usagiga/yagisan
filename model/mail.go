package model

type Mail struct {
	From string
	To   string // TODO: compatible with multiple Envelope To
	Body string
}
