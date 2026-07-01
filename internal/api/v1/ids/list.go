package ids

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/internal/api/ids"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// ListApplications for identifiers.
func (i *Identifier) ListApplications(ctx context.Context, _ *v1.ListApplicationsRequest) (*v1.ListApplicationsResponse, error) {
	apps := i.id.Applications()

	resp := &v1.ListApplicationsResponse{
		Meta:                  meta.CamelStrings(ctx, strings.Empty),
		GeneratorApplications: generatorApplications(apps.GeneratorApplications),
		MapperApplications:    mapperApplications(apps.MapperApplications),
		GeneratorKinds:        apps.GeneratorKinds,
		Limits: &v1.RequestLimits{
			GenerateCount: apps.Limits.GenerateCount,
			MapIds:        apps.Limits.MapIDs,
		},
	}

	return resp, nil
}

func generatorApplications(apps []*ids.GeneratorApplication) []*v1.GeneratorApplication {
	generatorApps := make([]*v1.GeneratorApplication, 0, len(apps))
	for _, app := range apps {
		generatorApps = append(generatorApps, &v1.GeneratorApplication{Name: app.Name, Kind: app.Kind})
	}

	return generatorApps
}

func mapperApplications(apps []*ids.MapperApplication) []*v1.MapperApplication {
	mapperApps := make([]*v1.MapperApplication, 0, len(apps))
	for _, app := range apps {
		mapperApps = append(mapperApps, &v1.MapperApplication{Name: app.Name})
	}

	return mapperApps
}
