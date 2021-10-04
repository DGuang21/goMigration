package migration

import "fmt"

/**
 * @Description:
 * @param tableName
 * @param version
 * @param fields
 * @auth: daguang
 */
func CreateTable(tableName string,version string,fields ...func(table *MigrationTable)) {
	if tableName == "" {
		panic("table can't be null")
	}
	fmt.Println(version)
	if len(fields) > 0 {
		for _,fieldFunc := range fields {
			a := MigrationTable{nil,nil}
			fieldFunc(&a)
			for _,vv := range a.result {
				fmt.Printf("111%+v\n",vv)
			}
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

/**
 * @Description:
 * @param tableName
 * @param version
 * @param fields
 * @auth: daguang
 */
func UpdateTable(tableName string,version string,fields ...func(table *MigrationTable)) {
	if tableName == "" {
		panic("table can't be null")
	}
	fmt.Println(version)
	if len(fields) > 0 {
		for _,fieldFunc := range fields {
			a := MigrationTable{nil,nil}
			fieldFunc(&a)
			for _,vv := range a.result {
				fmt.Printf("111%+v\n",vv)
			}
		}
	}
}

var (
	upList = []func(){}
	downList = []func(){}
)

/**
 * @Description:
 * @param up
 * @param down
 * @auth: daguang
 */
func RegisterMigration(up func(),down func()) {
	upList = append(upList, up)
	downList = append(downList, down)
}