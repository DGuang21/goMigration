package migration

type MigrationTable struct {
	*MigrationAttributes
	result []*MigrationAttributes
	tableName       string
	engine          string
	charset   		string
	collation		string
	temporary		bool
}

func (m *MigrationTable) Geometry(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "GEOMETRY",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) GeometryCollection(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "GEOMETRYCOLLECTION",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MultiPoint(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MULTIPOINT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MultiPolygon(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MULTIPOLYGON",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Point(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "POINT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Polygon(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "POLYGON",
	}
	return m.MigrationAttributes
}

// Comment 表注释
func (m *MigrationTable) Comment(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "comment",
	}
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

func (m *MigrationTable) Done() {
	m.appendLastResult()
}

func (m *MigrationTable) appendLastResult() {
	if m.MigrationAttributes != nil {
		m.result = append(m.result,m.MigrationAttributes)
	}
}