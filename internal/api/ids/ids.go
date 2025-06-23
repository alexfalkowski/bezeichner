package ids

import (
	"fmt"

	"github.com/alexfalkowski/bezeichner/internal/generator"
	"github.com/alexfalkowski/bezeichner/internal/mapper"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
)

// ErrNotFound for service.
var ErrNotFound = errors.New("not found")

// IsNotFound for service.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// NewIdentifier for the different transports.
func NewIdentifier(gc *generator.Config, mc *mapper.Config, gs generator.Generators) *Identifier {
	return &Identifier{generatorConfig: gc, mapperConfig: mc, generators: gs}
}

// Identifier for the different transports.
type Identifier struct {
	generatorConfig *generator.Config
	mapperConfig    *mapper.Config
	generators      generator.Generators
}

// Generate identifiers.
func (s *Identifier) Generate(ctx context.Context, application string, count uint64) ([]string, error) {
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

// Map identifiers.
func (s *Identifier) Map(ids []string) ([]string, error) {
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
