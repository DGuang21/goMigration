package migration

func (m *MigrationTable) Unique(name string, fields ...string) {
	m.table.uniqueName = append(m.table.uniqueName, name)
	m.table.uniqueFields = append(m.table.uniqueFields, fields)
}

func (m *MigrationTable) Index(indexName string, fields ...string) {
	m.table.indexName = append(m.table.indexName, indexName)
	m.table.indexFields = append(m.table.indexFields, fields)
}

func (m *MigrationTable) SpatialIndex(name string, fields ...string) {
	m.table.spatialName = append(m.table.spatialName, name)
	m.table.spatialFields = append(m.table.spatialFields, fields)
}

func (m *MigrationTable) Foreign(key string) *ForeignAttributes {
	if m.table.foreign == nil {
		m.table.foreign = &ForeignAttributes{}
	}
	m.table.foreign.result = append(m.table.foreign.result, &foreignAttribute{
		key: key,
	})
	return m.table.foreign
}

func (m *ForeignAttributes) References(table, field string) *ForeignAttributes {
	m.result[len(m.result)-1].referencesField = field
	m.result[len(m.result)-1].referencesTable = table
	return m
}

func (m *ForeignAttributes) OnUpdate() *ForeignAttributes {
	m.result[len(m.result)-1].onUpdate = true
	return m
}

func (m *ForeignAttributes) OnDelete() *ForeignAttributes {
	m.result[len(m.result)-1].onDelete = true
	return m
}

func (m *MigrationTable) RenameIndex(indexName, newName string) {

}

func (m *MigrationTable) DropPrimary(name string) {

}

func (m *MigrationTable) DropUnique(name string) {

}

func (m *MigrationTable) DropIndex(name string) {

}

func (m *MigrationTable) DropSpatialIndex(name string) {

}
