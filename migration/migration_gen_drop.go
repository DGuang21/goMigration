package migration

import "fmt"

func (m *MigrationTable) generateDropMigrationSQL() {
	for _,v := range m.table.result {
		fmt.Println(v)
	}
}
