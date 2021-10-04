package main

import "github.com/DGuang21/goMigration/migration"

type migrationCreateUserTableInterface struct {}

func init() {
	export := migrationCreateUserTableInterface{}
	migration.RegisterMigration(export.Up,export.Down)
}

func (m *migrationCreateUserTableInterface) Up() {
	migration.Create("asd1","20211004202802_c_up_date1", func(table *migration.MigrationTable) {
		table.String("555").NullTable(false).Comment("s")
		table.String("qq").Comment("s")
		table.Done()
	})
}

func (m *migrationCreateUserTableInterface) Down() {
	migration.DropIfExists("kos")
}

