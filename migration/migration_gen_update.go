package migration

import "fmt"

func (m *MigrationTable) generateUpdateMigrationSQL() {
	for _,v := range m.table.result {
		fmt.Printf("%+v\n",v)
	}
}
