package ids

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// GenerateIdentifiers for identifiers.
func (i *Identifier) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	ids, err := i.id.Generate(ctx, req.GetApplication(), req.GetCount())
	if err != nil {
		return nil, err
	}

	resp := &v1.GenerateIdentifiersResponse{
		Meta: meta.CamelStrings(ctx, strings.Empty),
		Ids:  ids,
	}

	return resp, nil
}
