package repository

import (
	"context"
	"fmt"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"
	"github.com/ognick/job-interview-playground/pkg/request"
)

type FavqsRepository struct {
}

func NewFavqsRepository() *FavqsRepository {
	return &FavqsRepository{}
}

func (repo *FavqsRepository) GetWisdom(ctx context.Context) (domain.Wisdom, error) {
	type dto struct {
		Quote struct {
			Body string `json:"body"`
		} `json:"quote"`
	}

	data, err := request.Get[dto](ctx, "https://favqs.com/api/qotd")
	if err != nil {
		return domain.Wisdom{}, fmt.Errorf("failed to get quote from favqs: %w", err)
	}

	wisdom := domain.Wisdom{
		Content: data.Quote.Body,
		Source:  "favqs",
	}

	return wisdom, nil
}
