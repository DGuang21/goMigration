package migration

func (m *MigrationTable) Unique(name string, fields ...string) {
	m.uniqueName = append(m.uniqueName, name)
	m.uniqueFields = append(m.uniqueFields, fields)
}

func (m *MigrationTable) Index(indexName string, fields ...string) {
	m.indexName = append(m.indexName, indexName)
	m.indexFields = append(m.indexFields, fields)
}

func (m *MigrationTable) SpatialIndex(name string, fields ...string) {
	m.spatialName = append(m.spatialName, name)
	m.spatialFields = append(m.spatialFields, fields)
}

func (m *MigrationTable) Foreign(key string) *ForeignAttributes {
	if m.foreign == nil {
		m.foreign = &ForeignAttributes{}
	}
	m.foreign.result = append(m.foreign.result, &foreignAttribute{
		key: key,
	})
	return m.foreign
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
