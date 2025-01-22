package wisdom

import (
	"context"
	"sync"
	"time"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"
	"github.com/ognick/job-interview-playground/pkg/logger"
)

const requestTimeout = 1 * time.Second

type repo interface {
	GetWisdom(ctx context.Context) (domain.Wisdom, error)
}

type Usecase struct {
	log           logger.Logger
	internalRepo  repo
	externalRepos []repo
}

func NewUsecase(
	log logger.Logger,
	internalRepo repo,
	externalRepos ...repo,
) *Usecase {
	return &Usecase{
		log:           log,
		internalRepo:  internalRepo,
		externalRepos: externalRepos,
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

	// Get wisdom from external repos concurrently
	for _, repo := range u.externalRepos {
		wg.Add(1)
		go func() {
			defer wg.Done()
			wisdom, err := repo.GetWisdom(ctx)
			out <- res{
				wisdom: wisdom,
				err:    err,
			}
		}()
	}
	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(out)
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
