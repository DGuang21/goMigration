package migration

type MigrationAttributes struct {
	field           string
	fieldType       string
	length          int
	precision       int
	defaultValue    interface{}
	canNull         bool
	comment         string
	maintenanceTime bool
	rename          string
	dropField       []string
	unique          bool
	references      string
	referencesOn    string
	isChange        bool
	primaryKey      bool
	autoIncrement   bool
}

/**
 * @Description: 字段重命名
 * @receiver m
 * @param field
 * @param newField
 * @auth: daguang
 */
func (m *MigrationTable) RenameField(field string, newField string) {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:  field,
		rename: newField,
	}
}

/**
 * @Description: 删除字段
 * @receiver m
 * @param fields
 * @auth: daguang
 */
func (m *MigrationTable) DropField(fields ...string) {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		dropField: fields,
	}
}

func (m *MigrationAttributes) Default(d interface{}) *MigrationAttributes {
	m.defaultValue = d
	return m
}

func (m *MigrationAttributes) NullTable(isNull bool) *MigrationAttributes {
	m.canNull = isNull
	return m
}

func (m *MigrationAttributes) Comment(c ...string) {
	comment := ""
	if len(c) != 0 {
		comment = c[0]
	}
	m.comment = comment
}

func (m *MigrationAttributes) Change() {
	m.isChange = true
}
