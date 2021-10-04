package main

import (
	"github.com/DGuang21/goMigration/migration"
)

type migrationCreatePowerTableInterface struct {}

func init() {
	export := migrationCreatePowerTableInterface{}
	migration.RegisterMigration(export.Up,export.Down)
}

func (m *migrationCreatePowerTableInterface) Up() {
	migration.Create("asd","20211004202802_c_up_date", func(table *migration.MigrationTable) {
		table.String("1").NullTable(false).Comment("s")
		table.BigIncrements("2")
		table.BigIncrements("3")
		table.BigIncrements("4")
		table.BigIncrements("5")
		table.Done()
	})
}

func (m *migrationCreatePowerTableInterface) Down() {
	migration.DropIfExists("kos")
}

