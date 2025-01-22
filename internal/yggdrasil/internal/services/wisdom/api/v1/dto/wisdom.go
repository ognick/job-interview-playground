package dto

import "github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"

type Wisdom struct {
	Message string `json:"message"`
	Source  string `json:"source"`
}

func NewWisdom(w domain.Wisdom) *Wisdom {
	return &Wisdom{
		Message: w.Content,
		Source:  w.Source,
	}
}
