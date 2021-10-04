package migration


func (m *MigrationTable) Foreign(field string) {

}

func (m *MigrationTable) Unique(name string) {

}

func (m *MigrationTable) Index(indexName string,fields ...string) {

}

func (m *MigrationTable) Primary(name ...string) {

}

func (m *MigrationTable) SpatialIndex(name string) {

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