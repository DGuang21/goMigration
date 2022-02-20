package migration

type migration struct{}

var (
	Migration = migration{}
)

/**
 * @Description:
 * @param tableName
 * @param version
 * @param fields
 * @auth: daguang
 */
func (m *migration) Create(tableName string, fields ...func(table *MigrationTable)) {
	if tableName == "" {
		panic("table can't be null")
	}
	if len(fields) > 0 {
		for _, fieldFunc := range fields {
			a := MigrationTable{
				table: &MigrationAttributes{
					tableName: tableName,
				},
			}
			fieldFunc(&a)
			a.generateCreateMigrationSQL()
		}
	}
}

/**
 * @Description:
 * @param tableName
 * @auth: daguang
 */
func (m *migration) DropIfExists(tableName string) {

}

func (m *migration) Drop(tableName string) {

}

/**
 * @Description:
 * @param tableName
 * @param version
 * @param fields
 * @auth: daguang
 */
func (m *migration) UpdateTable(tableName string, fields ...func(table *MigrationTable)) {
	if tableName == "" {
		panic("table can't be null")
	}
	if len(fields) > 0 {
		for _, fieldFunc := range fields {
			a := MigrationTable{}
			fieldFunc(&a)
		}
	}
}

func (m *migration) ReName(tableName string, newName string) {

}
