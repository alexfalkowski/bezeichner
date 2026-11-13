package ids

import (
	v1 "github.com/alexfalkowski/bezeichner/v2/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v3/context"
	"github.com/alexfalkowski/go-service/v3/meta"
)

// GenerateIdentifiers for identifiers.
func (i *Identifier) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	ids, err := i.id.Generate(ctx, req.GetApplication(), req.GetCount())
	if err != nil {
		return nil, err
	}

	resp := &v1.GenerateIdentifiersResponse{
		Meta: meta.Strings(ctx),
		Ids:  ids,
	}

	return resp, nil
}
