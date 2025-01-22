package repository

import (
	"context"
	"fmt"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"
	"github.com/ognick/job-interview-playground/pkg/request"
)

type QuotableRepository struct {
}

func NewQuotableRepository() *QuotableRepository {
	return &QuotableRepository{}
}

func (repo *QuotableRepository) GetWisdom(ctx context.Context) (domain.Wisdom, error) {
	type dto struct {
		Content string `json:"content"`
	}

	data, err := request.Get[dto](ctx, "http://api.quotable.io/random")
	if err != nil {
		return domain.Wisdom{}, fmt.Errorf("failed to get quote from quotable: %w", err)
	}

	wisdom := domain.Wisdom{
		Content: data.Content,
		Source:  "quotable",
	}

	return wisdom, nil
}
