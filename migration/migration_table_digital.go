package migration

func (m *MigrationTable) tinyInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "tinyInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) SmallInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "SMALLINT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Integer(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "INTEGER",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MediumInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MEDIUMINT",
	}
	return m.MigrationAttributes
}

/**
 * @Description: 相当于 BIGINT
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) BigInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "BIGINT",
	}
	return m.MigrationAttributes
}

/**
 * @Description:
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Boolean(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "tinyint",
	}
	return m.MigrationAttributes
}


func (m *MigrationTable) Decimal(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "DECIMAL",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Double(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "DOUBLE",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Float(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "FLOAT",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}


func (m *MigrationTable) unsignedBigInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedBigInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedDecimal(field string,length,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedBigInteger",
		length:     length,
		precision: precision,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedMediumInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedMediumInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedSmallInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedSmallInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedTinyInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedTinyInteger",
	}
	return m.MigrationAttributes
}