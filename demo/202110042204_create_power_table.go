package main

import (
	"github.com/DGuang21/goMigration/migration"
)

type migrationCreatePowerTableInterface struct {}

func init() {
	migration.RegisterMigration(&migrationCreatePowerTableInterface{},"20211004202802_c_up_date")
}

func (m *migrationCreatePowerTableInterface) Up() {
	migration.Create("asd",func(table *migration.MigrationTable) {
		table.String("1").NullTable(false).Comment("s")
		table.Timestamps().NullTable(false).Comment("s")
		table.Charset("utf8")
		table.Engine("InnoDB")
		table.Done()
	})
}

func (m *migrationCreatePowerTableInterface) Down() {
	migration.DropIfExists("kos")
}

