package migration


/**
 * @Description: 相当于为软删除添加一个可空的 deleted_at 字段
 * @receiver m
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) SoftDeletes() *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           "delete_at",
		fieldType:       "timestamp",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) Time(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "TIME",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) Timestamp(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "timestamp",
	})
	return m.MigrationAttributes
}

// 生成可以由框架自动维护的 create_at 与 update_at
func (m *MigrationTable) Timestamps() *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		fieldType:       "timestamps",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) Year(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "year",
	})
	return m.MigrationAttributes
}


func (m *MigrationTable) Date(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "DATE",
	})
	return m.MigrationAttributes
}

func (m *MigrationTable) DateTime(field string) *MigrationAttributes {
	m.MigrationAttributes.result = append(m.result,&MigrationAttribute{
		field:           field,
		fieldType:       "DATETIME",
	})
	return m.MigrationAttributes
}
