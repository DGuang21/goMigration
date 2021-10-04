package migration

type MigrationTable struct {
	*MigrationAttributes
	result []*MigrationAttributes
}

func (m *MigrationTable) BigIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "id",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) String(field string,length ...int) *MigrationAttributes {
	m.appendLastResult()
	l := 0
	if len(length) != 0 {
		l = length[0]
	}
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "string",
		length:          l,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Done() {
	m.appendLastResult()
}

func (m *MigrationTable) appendLastResult() {
	if m.MigrationAttributes != nil {
		m.result = append(m.result,m.MigrationAttributes)
	}
}