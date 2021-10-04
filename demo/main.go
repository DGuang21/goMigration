package main

import "github.com/DGuang21/goMigration/migration"

func main() {
	Migrate()
	migration.CleanMigration()
}
