package migration

import (
	"fmt"
	"log"

	"github.com/End-S/spooky_kingdom/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // required
	_ "github.com/golang-migrate/migrate/v4/source/file"       // required
)

func MigrateUp() {
	config := config.Get().Postgres
	fmt.Println("📥 Performing database schema migration")

	m, err := migrate.New(
		"file://db/migration/schema",
		fmt.Sprintf("postgres://%s:%s@%s:%s/spookydb", config.Username, config.Password, config.Host, config.Port),
	)

	if err != nil {
		fmt.Println("Failed to configure database schema migration")
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			fmt.Println("✅ No schema changes were required")
			return
		}
		fmt.Println("Failed to perform database schema migrate ")
		log.Fatal(err)
	}
}
