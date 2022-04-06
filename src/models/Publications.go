package models

import (
	"errors"
	"strings"
	"time"
)

type Publications struct {
	ID             uint64    `json:"id, omitempty"`
	Content        string    `json:"content, omitempty"`
	AuthorID       uint64    `json:"author_id, omitempty"`
	AuthorNickname string    `json:"author_nickname, omitempty"`
	LikeCount      uint64    `json:"like_count"`
	CreatedAt      time.Time `json:"created_at"`
}

func (p *Publications) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()

	return nil
}

func (p *Publications) validate() error {
	if p.Content == "" {
		return errors.New("Content is required")
	}

	return nil
}

func (p *Publications) format() {
	p.Content = strings.TrimSpace(p.Content)
}
