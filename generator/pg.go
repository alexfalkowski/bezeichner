package generator

import (
	"context"
	"fmt"
	"strconv"

	"github.com/linxGnu/mssqlx"
)

// PG ID generator.
type PG struct {
	name string
	db   *mssqlx.DBs
}

// Generate an ID using a sequence.
func (p *PG) Generate(ctx context.Context) (string, error) {
	var id int64

	row := p.db.QueryRowContext(ctx, fmt.Sprintf("SELECT nextval('%s')", p.name))
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}
