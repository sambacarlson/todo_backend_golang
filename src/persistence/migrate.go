package persistence

import "database/sql"

func CreateMigration(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	return err
}
