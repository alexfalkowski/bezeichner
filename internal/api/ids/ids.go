package ids

import (
	"fmt"

	"github.com/alexfalkowski/bezeichner/internal/generator"
	"github.com/alexfalkowski/bezeichner/internal/limits"
	"github.com/alexfalkowski/bezeichner/internal/mapper"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
)

// ErrNotFound indicates that a requested resource does not exist.
//
// In this package, it is used when:
//   - the requested generator application name cannot be found in configuration,
//   - the generator kind cannot be resolved from the generator registry,
//   - or the requested mapper application name cannot be found in configuration.
//
// Generate and Map may wrap ErrNotFound with the missing application, kind, or
// identifier value, so callers should classify it with errors.Is rather than
// direct equality.
var ErrNotFound = errors.New("not found")

// NewIdentifier constructs an Identifier domain service.
//
// It requires:
//   - generator configuration (to resolve an application by name),
//   - and a generator registry (to resolve a generator by application kind).
//
// Limits configuration is optional and defaults to the built-in domain limits.
// Mapper configuration is optional. When it is omitted, Map returns ErrNotFound.
func NewIdentifier(generator *generator.Config, mapper *mapper.Config, generators generator.Generators, limits *limits.Config) *Identifier {
	return &Identifier{generatorConfig: generator, mapperConfig: mapper, generators: generators, limits: limits}
}

// Identifier is the domain service that generates and maps identifiers.
//
// It is transport-agnostic and is intended to be used by multiple transports
// (for example gRPC and an HTTP gateway).
type Identifier struct {
	generatorConfig *generator.Config
	mapperConfig    *mapper.Config
	generators      generator.Generators
	limits          *limits.Config
}

// Generate returns count identifiers for the given application name.
//
// The application is resolved from generator configuration and selects the
// generator kind used to produce each identifier.
//
// A count of zero is valid and returns an empty slice when the application and
// generator kind can be resolved.
//
// Errors:
//   - ErrInvalidArgument if count exceeds the configured domain limit.
//   - ErrNotFound if the application name does not exist, or if the application
//     kind cannot be resolved to a generator.
func (s *Identifier) Generate(ctx context.Context, application string, count uint64) ([]string, error) {
	if count > s.limits.MaxGenerateCount() {
		return nil, ErrInvalidArgument
	}

	if s.generatorConfig == nil {
		return nil, ErrNotFound
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
		ids[i] = g.Generate(ctx, app)
	}

	return ids, nil
}

// Map classifies the provided ids into mapped and unmapped identifiers for the
// provided application.
//
// Mapping is configured via mapper application configuration. If mapper
// configuration is omitted or the application is missing, the operation fails.
// If an input identifier is missing from the application mapping, the returned
// unmapped list contains the identifier.
//
// An empty ids slice is valid and returns empty collections when mapper
// application configuration is present.
//
// Errors:
//   - ErrInvalidArgument if len(ids) exceeds the configured domain limit.
//   - ErrNotFound if mapper configuration is omitted, the application is
//     missing.
func (s *Identifier) Map(application string, ids []string) (mapper.Identifiers, []string, error) {
	if uint64(len(ids)) > s.limits.MaxMapIDs() {
		return nil, nil, ErrInvalidArgument
	}

	if s.mapperConfig == nil {
		return nil, nil, ErrNotFound
	}

	app := s.mapperConfig.Application(application)
	if app == nil {
		return nil, nil, fmt.Errorf("%s: %w", application, ErrNotFound)
	}

	mapped := make(mapper.Identifiers, len(ids))
	unmapped := make([]string, 0)
	for _, id := range ids {
		mid, ok := app.Identifiers[id]
		if !ok {
			unmapped = append(unmapped, id)
			continue
		}

		mapped[id] = mid
	}
	return mapped, unmapped, nil
}
