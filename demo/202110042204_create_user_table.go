package main

import "github.com/DGuang21/goMigration/migration"

type migrationCreateUserTableInterface struct {}

func init() {
	migration.RegisterMigration(&migrationCreateUserTableInterface{},"20211004202802_c_up_date1")
}

func (m *migrationCreateUserTableInterface) Up() {
	migration.DropIfExists("asd1")
	migration.Create("asd1", func(table *migration.MigrationTable) {
		table.String("555").NullTable(false).Comment("s")
		table.String("qq").Comment("s")
		table.Done()
	})
}

func (m *migrationCreateUserTableInterface) Down() {
	migration.DropIfExists("kos")
}

