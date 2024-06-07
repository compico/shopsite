package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateCategoryTable, downCreateCategoryTable)
}

func upCreateCategoryTable(ctx context.Context, tx *sql.Tx) error {
	q := `CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		parent_id INTEGER NOT NULL DEFAULT 0,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP NOT NULL DEFAULT now()
    );`

	_, err := tx.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func downCreateCategoryTable(ctx context.Context, tx *sql.Tx) error {
	q := `DROP TABLE IF EXISTS categories;`

	_, err := tx.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}
