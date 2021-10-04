package migration

type MigrationAttributes struct {
	field string
	fieldType string
	length int
	defaultValue interface{}
	canNull bool
	comment string
	maintenanceTime bool
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

func (m *MigrationAttributes) Timestamps() *MigrationAttributes {
	m.maintenanceTime = true
	return m
}