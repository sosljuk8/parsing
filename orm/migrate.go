package orm

import (
	"context"
	"log/slog"
)

// Migrate runs the migration tool
func Migrate() error {

	client := NewClient()

	defer func() {
		if err := client.Close(); err != nil {
			slog.Error("Fail to close database connection after migration", err)
		}
	}()

	return client.Schema.Create(context.Background())
}
