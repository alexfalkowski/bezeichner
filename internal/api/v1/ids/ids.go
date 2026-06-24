package ids

import (
	"github.com/alexfalkowski/bezeichner/internal/api/ids"
)

// NewIdentifier for identifiers.
func NewIdentifier(id *ids.Identifier) *Identifier {
	return &Identifier{id: id}
}

// Identifier for identifiers.
type Identifier struct {
	id *ids.Identifier
}
