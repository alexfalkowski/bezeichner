package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/mapper"
)

// ErrNotFound for service.
var ErrNotFound = errors.New("not found")

// IsNotFoundError for service.
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// NewService for the different transports.
func NewService(gc *generator.Config, mc *mapper.Config, gs generator.Generators) *Service {
	return &Service{generatorConfig: gc, mapperConfig: mc, generators: gs}
}

// Server for the different transports.
type Service struct {
	generatorConfig *generator.Config
	mapperConfig    *mapper.Config
	generators      generator.Generators
}

// GenerateIdentifiers for service.
func (s *Service) GenerateIdentifiers(ctx context.Context, application string, count uint64) ([]string, error) {
	if count == 0 {
		count = 1
	}

	app := s.generatorConfig.Application(application)
	if app == nil {
		return nil, fmt.Errorf("%s: %w", application, ErrNotFound)
	}

	g, err := s.generators.Generator(app.Kind)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", app.Kind, ErrNotFound, err)
	}

	ids := make([]string, count)
	for i := 0; i < len(ids); i++ {
		id, err := g.Generate(ctx, app)
		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	return ids, nil
}

// MapIdentifiers for service.
func (s *Service) MapIdentifiers(ids []string) ([]string, error) {
	mids := make([]string, len(ids))

	for i, id := range ids {
		mid, ok := s.mapperConfig.Identifiers[id]
		if !ok {
			return nil, fmt.Errorf("%s: %w", id, ErrNotFound)
		}

		mids[i] = mid
	}

	return mids, nil
}
