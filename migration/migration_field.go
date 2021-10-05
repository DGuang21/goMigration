package migration

type MigrationAttributes struct {
	tableName       string
	engine          string
	charset   		string
	collation		string
	temporary		bool
	result []*MigrationAttribute
}

type MigrationAttribute struct {
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
	unsigned		bool
}

/**
 * @Description: 字段重命名
 * @receiver m
 * @param field
 * @param newField
 * @auth: daguang
 */
func (m *MigrationTable) RenameField(field string, newField string) {
	m.result[len(m.result)-1].field = field
	m.result[len(m.result)-1].rename = newField
}

/**
 * @Description: 删除字段
 * @receiver m
 * @param fields
 * @auth: daguang
 */
func (m *MigrationTable) DropField(fields ...string) {
	m.result[len(m.result)-1].dropField = fields
}

func (m *MigrationAttributes) Default(d interface{}) *MigrationAttributes {
	m.result[len(m.result)-1].defaultValue = d
	return m
}

func (m *MigrationAttributes) NullTable(isNull bool) *MigrationAttributes {
	m.result[len(m.result)-1].canNull = isNull
	return m
}

func (m *MigrationAttributes) Comment(c ...string) {
	comment := ""
	if len(c) != 0 {
		comment = c[0]
	}
	m.result[len(m.result)-1].comment = comment
}

func (m *MigrationAttributes) Change() {
	m.result[len(m.result)-1].isChange = true
}
