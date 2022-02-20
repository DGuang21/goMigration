package migration

/**
 * @Description: 相当于 BLOB
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Binary(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "BLOB",
	})
	return m.table
}

/**
 * @Description: 相当于带有长度的 CHAR
 * @receiver m
 * @param field
 * @param length
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Char(field string, length ...int) *MigrationAttributes {
	l := 255
	if len(length) != 0 {
		l = length[0]
	}
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "char",
		length:    l,
	})
	return m.table
}

func (m *MigrationTable) Json(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "JSON",
	})
	return m.table
}

func (m *MigrationTable) LineString(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "LINESTRING",
	})
	return m.table
}

func (m *MigrationTable) LongText(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "LONGTEXT",
	})
	return m.table
}

func (m *MigrationTable) MediumText(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "MEDIUMTEXT",
	})
	return m.table
}

func (m *MigrationTable) MultiLineString(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "MULTILINESTRING",
	})
	return m.table
}

func (m *MigrationTable) Text(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "TEXT",
	})
	return m.table
}

func (m *MigrationTable) String(field string, length ...int) *MigrationAttributes {
	l := 255
	if len(length) != 0 {
		l = length[0]
	}
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "string",
		length:    l,
	})
	return m.table
}
