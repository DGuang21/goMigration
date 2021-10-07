package migration


func (m *MigrationTable) Foreign(field string) {

}

func (m *MigrationTable) Unique(name string,fields ...string) {
	m.uniqueName = append(m.uniqueName,name)
	m.uniqueFields = append(m.uniqueFields,fields)
}

func (m *MigrationTable) Index(indexName string,fields ...string) {
	m.indexName = append(m.indexName,indexName)
	m.indexFields = append(m.indexFields,fields)
}

func (m *MigrationTable) SpatialIndex(name string,fields ...string) {
	m.spatialName = append(m.spatialName,name)
	m.spatialFields = append(m.spatialFields,fields)
}

func (m *MigrationTable) RenameIndex(indexName,newName string) {

}

func (m *MigrationTable) DropPrimary(name string) {

}

func (m *MigrationTable) DropUnique(name string) {

}

func (m *MigrationTable) DropIndex(name string) {

}

func (m *MigrationTable) DropSpatialIndex(name string) {

}