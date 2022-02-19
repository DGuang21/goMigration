package migration

/**
 * @Description: 字段重命名
 * @receiver m
 * @param field
 * @param newField
 * @auth: daguang
 */
func (m *MigrationTable) RenameField(field string, newField string) {
	m.table.result[len(m.table.result)-1].field = field
	m.table.result[len(m.table.result)-1].rename = newField
}

/**
 * @Description: 删除字段
 * @receiver m
 * @param fields
 * @auth: daguang
 */
func (m *MigrationTable) DropField(fields ...string) {
	m.table.result[len(m.table.result)-1].dropField = fields
}

func (m *MigrationAttributes) Default(d interface{}) *MigrationAttributes {
	if len(m.result) == 0 {
		panic("没有指定字段类型")
	}
	m.result[len(m.result)-1].defaultValue = d
	return m
}

func (m *MigrationAttributes) DefaultNow() *MigrationAttributes {
	if len(m.result) == 0 {
		panic("没有指定字段类型")
	}
	m.result[len(m.result)-1].defaultValue = "CURRENT_TIMESTAMP"
	return m
}

func (m *MigrationAttributes) NullTable(isNull bool) *MigrationAttributes {
	if len(m.result) == 0 {
		panic("没有指定字段类型")
	}
	m.result[len(m.result)-1].canNull = isNull
	return m
}

func (m *MigrationAttributes) Comment(c ...string) {
	if len(m.result) == 0 {
		panic("没有指定字段类型")
	}
	comment := ""
	if len(c) != 0 {
		comment = c[0]
	}
	m.result[len(m.result)-1].comment = comment
}

func (m *MigrationAttributes) Change() {
	m.result[len(m.result)-1].isChange = true
}

// Index 添加一个单独的 index
func (m *MigrationAttributes) Index(indexName string) *MigrationAttributes {
	m.indexFields = append(m.indexFields,[]string{m.result[len(m.result)-1].field})
	m.indexName = append(m.indexName,indexName)
	return m
}