package migration

type migrationer interface {

	Up()

	Down()

}