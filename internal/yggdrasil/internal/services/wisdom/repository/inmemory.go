package repository

import (
	"context"
	"math/rand"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"
)

type InmemoryRepository struct {
	words []string
}

func NewInmemoryRepository() *InmemoryRepository {
	return &InmemoryRepository{
		words: []string{
			"Don't communicate by sharing memory, share memory by communicating",
			"Concurrency is not parallelism",
			"Channels orchestrate; mutexes serialize",
			"The bigger the interface, the weaker the abstraction",
			"Make the zero value useful",
			"interface{} says nothing",
			"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite",
			"A little copying is better than a little dependency",
			"Syscall must always be guarded with build tags",
			"Cgo must always be guarded with build tags",
			"Errors are values",
			"Don't just check errors, handle them gracefully",
			"Design the architecture, name the components, document the details",
			"Documentation is for users",
		},
	}
}

func (repo *InmemoryRepository) GetWisdom(_ context.Context) (domain.Wisdom, error) {
	size := len(repo.words)
	pos := rand.Intn(size)
	wisdom := domain.Wisdom{
		Content: repo.words[pos],
	}
	return wisdom, nil
}
