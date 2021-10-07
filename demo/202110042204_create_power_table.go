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
		table.String("varchar").NullTable(false).Comment("s")
		table.Timestamps().NullTable(false).Comment("s")
		table.SoftDeletes()
		table.Time("time").Default("1111").Comment("2222")
		table.Timestamp("Timestamp").DefaultNow()
		table.Year("Year")
		table.Date("Date")
		table.DateTime("DateTime")
		table.TinyInteger("TinyInteger")
		table.SmallInteger("SmallInteger")
		table.MediumInteger("MediumInteger")
		table.Integer("Integer")
		table.BigInteger("BigInteger")
		table.Boolean("Boolean").Default(true)
		table.Decimal("Decimal",10,2)
		table.Double("Double",10,2)
		table.UnsignedBigInteger("UnsignedBigInteger")
		table.Float("Float",10,2)
		table.UnsignedDecimal("UnsignedDecimal",10,2)
		table.UnsignedInteger("UnsignedInteger")
		table.UnsignedMediumInteger("UnsignedMediumInteger")
		table.UnsignedSmallInteger("UnsignedSmallInteger")
		table.UnsignedTinyInteger("UnsignedTinyInteger")
		table.Char("Char")
		table.Json("Json")
		table.LineString("LineString")
		table.LongText("LongText")
		table.MediumText("MediumText")
		table.MultiLineString("MultiLineString")
		table.Text("Text")
		table.String("stringss").Default("asdasd")
		table.Charset("utf8")
		table.Engine("InnoDB")
	})
}

func (m *migrationCreatePowerTableInterface) Down() {
	migration.DropIfExists("kos")
}

