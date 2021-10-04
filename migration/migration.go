package migration

import "fmt"

/**
 * @Description:
 * @param tableName
 * @param version
 * @param fields
 * @auth: daguang
 */
func Create(tableName string,version string,fields ...func(table *MigrationTable)) {
	if tableName == "" {
		panic("table can't be null")
	}
	if len(fields) > 0 {
		for _,fieldFunc := range fields {
			a := MigrationTable{nil,nil}
			fieldFunc(&a)
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

func ReName(tableName string,newName string) {

}