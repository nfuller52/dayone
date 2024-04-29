//go:build dev

package cli

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

func RunCLI(args []string) {
	if len(args) == 1 {
		os.Exit(0)
	}

	switch args[1] {
	case "help":
		fmt.Println("🤖 dayone CLI helper")
	case "version":
		fmt.Println("🤖0.1.0")
	case "migrate":
		migrate()
	case "migrate:create":
		migrateCreate(args[2])
	case "migrate:rollback":
		fmt.Println("🤖 dayone CLI helper: migrate:rollback")
	case "migrate:status":
		fmt.Println("🤖 dayone CLI helper: migrate:status")
	default:
		fmt.Println("🤖 dayone CLI helper: command not found")
	}
}

func migrate() {
	fmt.Println("--> Migrating the database")

	// - Read all the files in the application/database/migrations folder
	// - Filter out any files that don't end in .go
	// - Split on file name _
	// - Sort by the first part of the file name (should be datetime)
	// - Check the db `dayone.app_migrations` table for the last migration
	//   - If there is no last migration, run all the migrations
	//   - If there is a last migration, run all the migrations after that
	//   - If there is a last migration, and the last migration is the last migration in the list, noop
	// - Run the method `up` in each migration file, in order
	//   - Create a record in the `dayone.app_migrations` table for each migration
	//   - If there is an error, rollback all the migrations (so run the migrations in a transaction)
	//
}

func migrateCreate(migrationTitle string) {
	currentDateTime := time.Now().Format("20060102150405")
	migrationTitle = strings.ToLower(migrationTitle)
	migrationTitle = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(migrationTitle, "_")
	fileName := fmt.Sprintf("%s_%s.go", currentDateTime, migrationTitle)

	filePath := path.Join("application", "database", "migrations", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create migration file:", err)
		return
	}
	defer file.Close()

	migrationCode := fmt.Sprintf(`
package migrations

import (
	"database/sql"
)

migrationVersion := "%s"

func Up() error {
	// SQL code to apply the migration
}

func Down() error {
	// SQL code to rollback the migration
}
`, currentDateTime)

	_, err = file.WriteString(migrationCode)
	if err != nil {
		fmt.Println("⚠️ Failed to write migration code to file:", err)
		return
	}

	fmt.Println("✅ New migration file created:", filePath)
}
