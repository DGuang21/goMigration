package migration

/**
 * @Description:
 * @param tableName
 * @param version
 * @param fields
 * @auth: daguang
 */
func Create(tableName string, fields ...func(table *MigrationTable)) {
	if tableName == "" {
		panic("table can't be null")
	}
	if len(fields) > 0 {
		for _, fieldFunc := range fields {
			a := MigrationTable{
				table: &MigrationAttributes{
					tableName:     tableName,
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
func DropIfExists(tableName string) {

}

func Drop(tableName string) {

}

/**
 * @Description:
 * @param tableName
 * @param version
 * @param fields
 * @auth: daguang
 */
func UpdateTable(tableName string, fields ...func(table *MigrationTable)) {
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

func ReName(tableName string, newName string) {

}
