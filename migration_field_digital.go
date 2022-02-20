package migration

func (m *MigrationTable) TinyInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "TINYINT",
	})
	return m.table
}

func (m *MigrationTable) SmallInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "SMALLINT",
	})
	return m.table
}

func (m *MigrationTable) MediumInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "MEDIUMINT",
	})
	return m.table
}

func (m *MigrationTable) Integer(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "INTEGER",
	})
	return m.table
}

/**
 * @Description: 相当于 BIGINT
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) BigInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "BIGINT",
	})
	return m.table
}

/**
 * @Description:
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Boolean(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "tinyint",
	})
	return m.table
}

func (m *MigrationTable) Decimal(field string, length int, precision int) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "DECIMAL",
		precision: precision,
		length:    length,
	})
	return m.table
}

func (m *MigrationTable) Double(field string, length int, precision int) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "DOUBLE",
		precision: precision,
		length:    length,
	})
	return m.table
}

func (m *MigrationTable) Float(field string, length int, precision int) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "FLOAT",
		precision: precision,
		length:    length,
	})
	return m.table
}

func (m *MigrationTable) UnsignedBigInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "BIGINT",
		unsigned:  true,
	})
	return m.table
}

func (m *MigrationTable) UnsignedDecimal(field string, length, precision int) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "DECIMAL",
		length:    length,
		precision: precision,
		unsigned:  true,
	})
	return m.table
}

func (m *MigrationTable) UnsignedInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "INTEGER",
		unsigned:  true,
	})
	return m.table
}

func (m *MigrationTable) UnsignedMediumInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "MEDIUMINT",
		unsigned:  true,
	})
	return m.table
}

func (m *MigrationTable) UnsignedSmallInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "SMALLINT",
		unsigned:  true,
	})
	return m.table
}

func (m *MigrationTable) UnsignedTinyInteger(field string) *MigrationAttributes {
	m.table.result = append(m.table.result, &MigrationAttribute{
		field:     field,
		fieldType: "TINYINT",
		unsigned:  true,
	})
	return m.table
}
