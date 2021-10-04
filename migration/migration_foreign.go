package migration

type MigrationForegion struct {
	*MigrationAttributes
}

func (m *MigrationForegion) References(field string) *MigrationForegion {
	m.MigrationAttributes.references = field
	return m
}

func (m *MigrationForegion) On(table string) {
	m.MigrationAttributes.referencesOn = table
}