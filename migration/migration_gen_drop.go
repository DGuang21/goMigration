package migration

import "fmt"

func (m *MigrationTable) generateDropMigrationSQL() {
	for _,v := range m.result {
		fmt.Println(m.engine)
		fmt.Println(m.charset)
		fmt.Println(m.collation)
		fmt.Printf("%+v\n",v)
	}
}
