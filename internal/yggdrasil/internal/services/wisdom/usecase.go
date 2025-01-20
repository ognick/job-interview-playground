package wisdom

import (
	"context"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"
	"github.com/ognick/job-interview-playground/pkg/logger"
)

type repo interface {
	GetWisdom(ctx context.Context) (domain.Wisdom, error)
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
	return u.internalRepo.GetWisdom(ctx)
}
