package migration


/**
 * @Description: 相当于 BLOB
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Binary(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "BLOB",
	}
	return m.MigrationAttributes
}

/**
 * @Description: 相当于带有长度的 CHAR
 * @receiver m
 * @param field
 * @param length
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Char(field string,length ...int) *MigrationAttributes {
	m.appendLastResult()
	l := 0
	if len(length) != 0 {
		l = length[0]
	}
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "char",
		length: l,
	}
	return m.MigrationAttributes
}


func (m *MigrationTable) Json(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "JSON",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Jsonb(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "JSONB",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) LineString(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "LINESTRING",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) LongText(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "LONGTEXT",
	}
	return m.MigrationAttributes
}


func (m *MigrationTable) MediumText(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MEDIUMTEXT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MultiLineString(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MULTILINESTRING",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Text(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "TEXT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) UUID(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "uuid",
	}
	return m.MigrationAttributes
}


func (m *MigrationTable) String(field string,length ...int) *MigrationAttributes {
	m.appendLastResult()
	l := 0
	if len(length) != 0 {
		l = length[0]
	}
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "string",
		length:          l,
	}
	return m.MigrationAttributes
}