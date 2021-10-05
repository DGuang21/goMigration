package migration

/**
 * @Description: 递增 ID（主键），相当于「UNSIGNED BIG INTEGER」
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) BigIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:         field,
		fieldType:     "BIGINT",
		primaryKey:    true,
		autoIncrement: true,
		unsigned:      true,
	}
	return m.MigrationAttributes
}

/**
 * @Description: 递增的 ID (主键)，相当于「UNSIGNED INTEGER」
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Increments(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:         field,
		fieldType:     "INTEGER",
		primaryKey:    true,
		autoIncrement: true,
		unsigned:      true,
	}
	return m.MigrationAttributes
}

/**
 * @Description: 递增 ID (主键) ，相当于「UNSIGNED MEDIUM INTEGER」
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) MediumIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:         field,
		fieldType:     "MEDIUMINT",
		primaryKey:    true,
		autoIncrement: true,
		unsigned:      true,
	}
	return m.MigrationAttributes
}

/**
 * @Description: 递增 ID (主键) ，相当于「UNSIGNED SMALL INTEGER」
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) SmallIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:         field,
		fieldType:     "SMALLINT",
		primaryKey:    true,
		autoIncrement: true,
		unsigned:      true,
	}
	return m.MigrationAttributes
}
