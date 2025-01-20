package dto

import "github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"

type Wisdom struct {
	Message string `json:"message"`
}

func NewWisdom(w domain.Wisdom) *Wisdom {
	return &Wisdom{
		Message: w.Content,
	}
}
