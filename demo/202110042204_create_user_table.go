package main

import "github.com/DGuang21/goMigration/migration"

type migrationCreateUserTableInterface struct {}

func init() {
	export := migrationCreateUserTableInterface{}
	migration.RegisterMigration(export.Up,export.Down)
}

func (m *migrationCreateUserTableInterface) Up() {
	migration.Create("asd","20211004202802_c_up_date", func(table *migration.MigrationTable) {
		table.String("1").NullTable(false).Comment("s")
		table.Timestamps()
		table.Done()
	})
}

func (m *migrationCreateUserTableInterface) Down() {
	migration.DropIfExists("kos")
}

