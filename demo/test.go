package demo

import "github.com/DGuang21/goMigration/migration"

type migrationCUpDateer struct {}

func init() {
	export := migrationCUpDateer{}
	migration.RegisterMigration(export.Up,export.Down)
}

func (m *migrationCUpDateer) Up() {
	migration.CreateTable("asd","20211004202802_c_up_date", func(table *migration.MigrationTable) {
		table.String("1").NullTable(false).Comment("s")
		table.BigIncrements("2").Timestamps().Comment()
		table.BigIncrements("3").Timestamps()
		table.BigIncrements("4").Timestamps()
		table.BigIncrements("5").Timestamps()
		table.Done()
	})
}

func (m *migrationCUpDateer) Down() {
	migration.DropIfExists("kos")
}

