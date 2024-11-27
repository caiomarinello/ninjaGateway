package repositories

import (
	"context"
	"database/sql"
)

func DeleteSession(db *sql.DB, sessionID string) error {
	ctx := context.Background()
	_, err := db.ExecContext(ctx, "DELETE FROM sessions WHERE id = ?", sessionID)
	return err
}
