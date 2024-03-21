package migrations

import (
	"database/sql"
	"embed"
	_ "github.com/jackc/pgx/v5/stdlib"
	migrate "github.com/rubenv/sql-migrate"
	"io/fs"
	"net/http"
)

var (
	//go:embed schema/*.sql
	migrations embed.FS
)

func Apply(dsn string) (applied int, err error) {
	subFS, _ := fs.Sub(migrations, "schema")
	staticFS := http.FS(subFS)
	var migrationSource = &migrate.HttpFileSystemMigrationSource{
		FileSystem: staticFS,
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	return migrate.Exec(db, "postgres", migrationSource, migrate.Up)
}
