package ids

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// MapIdentifiers for identifiers.
func (i *Identifier) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	mapped, unmapped, err := i.id.Map(req.GetApplication(), req.GetIds())
	if err != nil {
		return nil, err
	}

	resp := &v1.MapIdentifiersResponse{
		Meta:     meta.CamelStrings(ctx, strings.Empty),
		Mapped:   mapped,
		Unmapped: unmapped,
	}

	return resp, nil
}
