package wisdom

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"
	"github.com/ognick/job-interview-playground/pkg/logger"
	"github.com/ognick/job-interview-playground/pkg/request"
)

const requestTimeout = 1 * time.Second

type repo interface {
	GetWisdom(ctx context.Context) (domain.Wisdom, error)
}

func getFavqsQuote(ctx context.Context) (string, error) {
	type dto struct {
		Quote struct {
			Body string `json:"body"`
		} `json:"quote"`
	}

	data, err := request.Get[dto](ctx, "https://favqs.com/api/qotd")
	if err != nil {
		return "", fmt.Errorf("failed to get quote from favqs: %w", err)
	}

	return data.Quote.Body, nil
}

func getQuotableQuote(ctx context.Context) (string, error) {
	type dto struct {
		Content string `json:"content"`
	}

	data, err := request.Get[dto](ctx, "http://api.quotable.io/random")
	if err != nil {
		return "", fmt.Errorf("failed to get quote from favqs: %w", err)
	}

	return data.Content, nil
}

func getFavqsWisdom(ctx context.Context) (string, error) {
	type dto struct {
		Quote struct {
			Body string `json:"body"`
		} `json:"quote"`
	}

	data, err := request.Get[dto](ctx, "https://favqs.com/api/qotd")
	if err != nil {
		return "", fmt.Errorf("failed to get quote from favqs: %w", err)
	}

	return data.Quote.Body, nil
}

type Usecase struct {
	log          logger.Logger
	internalRepo repo
}

func NewUsecase(
	log logger.Logger,
	internalRepo repo,
) *Usecase {
	return &Usecase{
		log:          log,
		internalRepo: internalRepo,
	}
}

func (u *Usecase) GetWisdom(ctx context.Context) (domain.Wisdom, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	type res struct {
		wisdom domain.Wisdom
		err    error
	}
	out := make(chan res)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(out)
	}()
	// Run the goroutines to fetch data from the quote providers1
	go func() {
		defer wg.Done()
		quote, err := getFavqsWisdom(ctx)
		out <- res{
			wisdom: domain.Wisdom{
				Content: quote,
				Source:  "favqs",
			},
			err: err,
		}
	}()
	// Run the goroutines to fetch data from the quote providers2
	go func() {
		defer wg.Done()
		quote, err := getQuotableQuote(ctx)
		out <- res{
			wisdom: domain.Wisdom{
				Content: quote,
				Source:  "quotable",
			},
			err: err,
		}
	}()
	// Return the first successful result
	for r := range out {
		if r.err == nil {
			return r.wisdom, nil
		}
	}
	// Return the error if all the results are failed
	return u.internalRepo.GetWisdom(ctx)
}
