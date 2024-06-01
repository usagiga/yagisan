package model

import "net/url"

// Address represents Discord Channel where messages are send to
type Address struct {
	ID         string
	WebhookURL url.URL
}
