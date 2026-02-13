package ids

import (
	"fmt"

	"github.com/alexfalkowski/bezeichner/internal/generator"
	"github.com/alexfalkowski/bezeichner/internal/mapper"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
)

// ErrNotFound indicates that a requested resource does not exist.
//
// In this package, it is used when:
//   - the requested generator application name cannot be found in configuration,
//   - the generator kind cannot be resolved from the generator registry,
//   - or an input identifier does not have a configured mapping.
var ErrNotFound = errors.New("not found")

// IsNotFound reports whether err is (or wraps) ErrNotFound.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// NewIdentifier constructs an Identifier domain service.
//
// It requires:
//   - generator configuration (to resolve an application by name),
//   - mapper configuration (to map identifiers),
//   - and a generator registry (to resolve a generator by application kind).
func NewIdentifier(gc *generator.Config, mc *mapper.Config, gs generator.Generators) *Identifier {
	return &Identifier{generatorConfig: gc, mapperConfig: mc, generators: gs}
}

// Identifier is the domain service that generates and maps identifiers.
//
// It is transport-agnostic and is intended to be used by multiple transports
// (for example gRPC and an HTTP gateway).
type Identifier struct {
	generatorConfig *generator.Config
	mapperConfig    *mapper.Config
	generators      generator.Generators
}

// Generate returns count identifiers for the given application name.
//
// The application is resolved from generator configuration and selects the
// generator kind used to produce each identifier.
//
// Errors:
//   - ErrInvalidArgument if count exceeds the configured limit.
//   - ErrNotFound if the application name does not exist, or if the application
//     kind cannot be resolved to a generator.
func (s *Identifier) Generate(ctx context.Context, application string, count uint64) ([]string, error) {
	if count > maxGenerateCount {
		return nil, ErrInvalidArgument
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
	for i := range ids {
		id, err := g.Generate(ctx, app)
		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	return ids, nil
}

// Map returns the mapped identifiers for the provided ids in the same order.
//
// Mapping is configured via mapper configuration. If any input identifier is
// missing from the mapping table, the operation fails.
//
// Errors:
//   - ErrInvalidArgument if len(ids) exceeds the configured limit.
//   - ErrNotFound if any input identifier does not have a configured mapping.
func (s *Identifier) Map(ids []string) ([]string, error) {
	if len(ids) > maxMapIDs {
		return nil, ErrInvalidArgument
	}

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
