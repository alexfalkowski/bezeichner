package generator

import (
	"fmt"
	"strconv"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/strings"
	"github.com/linxGnu/mssqlx"
)

// PG generator.
type PG struct {
	db *mssqlx.DBs
}

// Generate an ID using a sequence.
func (p *PG) Generate(ctx context.Context, app *Application) (string, error) {
	var id int64

	row := p.db.QueryRowContext(ctx, fmt.Sprintf("SELECT nextval('%s')", app.Name))
	if err := row.Scan(&id); err != nil {
		return strings.Empty, err
	}

	return app.ID(strconv.FormatInt(id, 10)), nil
}
