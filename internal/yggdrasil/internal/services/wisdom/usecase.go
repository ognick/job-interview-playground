package wisdom

import (
	"context"
	"sync"
	"time"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"
	"github.com/ognick/job-interview-playground/pkg/logger"
	"github.com/ognick/job-interview-playground/pkg/quotes/favqs"
	"github.com/ognick/job-interview-playground/pkg/quotes/quotable"
)

const requestTimeout = 1 * time.Second

type repo interface {
	GetWisdom(ctx context.Context) (domain.Wisdom, error)
}

type Usecase struct {
	log            logger.Logger
	internalRepo   repo
	quoteProvider1 *favqs.FavqsQuoteProvider
	quoteProvider2 *quotable.QuotableQuoteProvider
}

func NewUsecase(
	log logger.Logger,
	internalRepo repo,
) *Usecase {
	return &Usecase{
		log:            log,
		internalRepo:   internalRepo,
		quoteProvider1: favqs.NewFavqsQuoteProvider(),
		quoteProvider2: quotable.NewQuotableQuoteProvider(),
	}
}

func (u *Usecase) GetWisdom(c context.Context) (domain.Wisdom, error) {
	ctx, cancel := context.WithTimeout(c, requestTimeout)
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
		data, err := u.quoteProvider1.GetData(ctx)
		out <- res{
			wisdom: domain.Wisdom{
				Content: data.Quote.Body,
				Source:  "Favqs",
			},
			err: err,
		}
	}()
	// Run the goroutines to fetch data from the quote providers2
	go func() {
		defer wg.Done()
		data, err := u.quoteProvider2.GetData(ctx)
		out <- res{
			wisdom: domain.Wisdom{
				Content: data.Content,
				Source:  "Quotable",
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
