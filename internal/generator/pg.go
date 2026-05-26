package generator

import (
	"strconv"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/database/sql"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// PG generator.
type PG struct {
	db *sql.DBs
}

// Generate an ID using a sequence.
func (p *PG) Generate(ctx context.Context, app *Application) (string, error) {
	if p.db == nil {
		return strings.Empty, ErrUnavailable
	}

	var id int64

	// nextval advances sequence state, so it must run against the master pool.
	row := p.db.QueryRowContextOnMaster(ctx, "SELECT nextval($1::regclass)", app.Name)
	if err := row.Scan(&id); err != nil {
		return strings.Empty, err
	}

	return strconv.FormatInt(id, 10), nil
}
