package migration

type MigrationTable struct {
	*MigrationAttributes
}

func (m *MigrationTable) Geometry(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "GEOMETRY",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) GeometryCollection(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "GEOMETRYCOLLECTION",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) MultiPoint(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "MULTIPOINT",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) MultiPolygon(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "MULTIPOLYGON",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) Point(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "POINT",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) Polygon(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "POLYGON",
	})
	return m.MigrationAttributes
}

// Comment 表注释
func (m *MigrationTable) Comment(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "comment",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) Engine(engine string) {
	m.engine = engine
}

func (m *MigrationTable) Charset(charset string) {
	m.charset = charset
}

func (m *MigrationTable) Collation(collation string) {
	m.collation = collation
}

func (m *MigrationTable) Temporary() {
	m.temporary = true
}