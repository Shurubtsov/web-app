package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: matching entry not found")

type Snippet struct {
	ID      int       `json:"id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Expires time.Time `json:"expires,omitempty"`
}
