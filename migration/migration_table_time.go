package migration


/**
 * @Description: 相当于为软删除添加一个可空的 deleted_at 字段
 * @receiver m
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) SoftDeletes() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           "delete_at",
		fieldType:       "timestamp",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Time(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "TIME",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) TimeTz(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timeTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Timestamp(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timestamp",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) TimestampTz(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timestampTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Timestamps() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		fieldType:       "timestamps",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) TimestampsTz() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           "",
		fieldType:       "timestampsTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) TinyIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "tinyIncrements",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Year(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "year",
	}
	return m.MigrationAttributes
}


func (m *MigrationTable) Date(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "DATE",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) DateTime(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "DATETIME",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) DateTimeTz(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timestamp",
	}
	return m.MigrationAttributes
}
