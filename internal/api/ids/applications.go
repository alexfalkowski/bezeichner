package ids

import "slices"

// Applications contains safe application discovery data.
type Applications struct {
	Limits                *Limits
	GeneratorApplications []*GeneratorApplication
	MapperApplications    []*MapperApplication
	GeneratorKinds        []string
}

// GeneratorApplication describes a configured generator application.
type GeneratorApplication struct {
	Name string
	Kind string
}

// MapperApplication describes a configured mapper application.
type MapperApplication struct {
	Name string
}

// Limits contains effective per-request item-count limits.
type Limits struct {
	GenerateCount uint64
	MapIDs        uint64
}

// Applications returns configured application names and safe capability data.
func (s *Identifier) Applications() *Applications {
	return &Applications{
		GeneratorApplications: s.generatorApplications(),
		MapperApplications:    s.mapperApplications(),
		GeneratorKinds:        s.generatorKinds(),
		Limits: &Limits{
			GenerateCount: s.limits.MaxGenerateCount(),
			MapIDs:        s.limits.MaxMapIDs(),
		},
	}
}

func (s *Identifier) generatorApplications() []*GeneratorApplication {
	if s.generatorConfig == nil {
		return nil
	}

	apps := make([]*GeneratorApplication, 0, len(s.generatorConfig.Applications))
	for _, app := range s.generatorConfig.Applications {
		apps = append(apps, &GeneratorApplication{Name: app.Name, Kind: app.Kind})
	}

	return apps
}

func (s *Identifier) mapperApplications() []*MapperApplication {
	if s.mapperConfig == nil {
		return nil
	}

	apps := make([]*MapperApplication, 0, len(s.mapperConfig.Applications))
	for _, app := range s.mapperConfig.Applications {
		apps = append(apps, &MapperApplication{Name: app.Name})
	}

	return apps
}

func (s *Identifier) generatorKinds() []string {
	kinds := make([]string, 0, len(s.generators))
	for kind := range s.generators {
		kinds = append(kinds, kind)
	}

	slices.Sort(kinds)

	return kinds
}
