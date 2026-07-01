package ids

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/internal/api/ids"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// MapIdentifiers for identifiers.
func (i *Identifier) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	ids, err := i.id.Map(req.GetApplication(), req.GetIds())
	if err != nil {
		return nil, err
	}

	resp := &v1.MapIdentifiersResponse{
		Meta: meta.CamelStrings(ctx, strings.Empty),
		Ids:  mappedIdentifiers(ids),
	}

	return resp, nil
}

func mappedIdentifiers(ids []*ids.MappedIdentifier) []*v1.MappedIdentifier {
	results := make([]*v1.MappedIdentifier, 0, len(ids))
	for _, id := range ids {
		results = append(results, &v1.MappedIdentifier{Id: id.ID, Mapped: id.Mapped})
	}

	return results
}
