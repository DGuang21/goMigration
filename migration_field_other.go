package migration

func (m *MigrationTable) Geometry(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "GEOMETRY",
	})
	return m.table
}

func (m *MigrationTable) GeometryCollection(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "GEOMETRYCOLLECTION",
	})
	return m.table
}

func (m *MigrationTable) MultiPoint(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "MULTIPOINT",
	})
	return m.table
}

func (m *MigrationTable) MultiPolygon(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "MULTIPOLYGON",
	})
	return m.table
}

func (m *MigrationTable) Point(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "POINT",
	})
	return m.table
}

func (m *MigrationTable) Polygon(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "POLYGON",
	})
	return m.table
}

// Comment 表注释
func (m *MigrationTable) Comment(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "comment",
	})
	return m.table
}

func (m *MigrationTable) Engine(engine string) {
	m.table.engine = engine
}

func (m *MigrationTable) Charset(charset string) {
	m.table.charset = charset
}

func (m *MigrationTable) Collation(collation string) {
	m.table.collation = collation
}

func (m *MigrationTable) Temporary() {
	m.table.temporary = true
}
