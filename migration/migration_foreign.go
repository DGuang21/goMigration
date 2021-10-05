package migration

type MigrationForegion struct {
	*MigrationAttributes
}

func (m *MigrationForegion) References(field string) *MigrationForegion {
	return m
}

func (m *MigrationForegion) On(table string) {

}