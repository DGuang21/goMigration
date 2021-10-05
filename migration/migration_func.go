package migration

import "database/sql"

var (
	upMap       = map[string]func(){}
	downMap     = map[string]func(){}
	versionList = []string{}
	tx *sql.Tx
)

func RegisterSqlDrive(t *sql.Tx) {
	tx = t
}

func testConn() {
	if tx == nil {
		panic("sql drive can't nil")
	}
}

// RegisterMigration
// All migration methods are registered with Slice during initialization
// 初始化时将全部迁移的 UP/DOWN 方法添加到 slice 和对应的 map 中
func RegisterMigration(migrationFunc migrationer,version string) {
	upMap[version] = migrationFunc.Up
	downMap[version] = migrationFunc.Down
	versionList = append(versionList, version)
}

// CleanMigration
// Clean up funcs that are permanently up/down in memory
// 清理在初始化时常驻在内存中的 up/down 方法
func CleanMigration() {
	upMap = nil
	downMap = nil
	versionList = nil
}

// Migrate 运行迁移
func Migrate(force ...bool) {
	//testConn()
	// diff
	for _, version := range versionList {
		vFunc, ok := upMap[version]
		if ok {
			if len(force) != 0 {
				if force[0] {
					// 强制执行(已经存在的表
				}
			}
			vFunc()
		}
	}
}

// RollBack 迁移回滚,如果不指定步数则默认为 1
func RollBack(step ...int) {
	testConn()
}

// RollBackAndMigrate 先执行迁移回滚然后重新执行迁移
func RollBackAndMigrate(step ...int) {
	testConn()
}

// Fresh 清空所有数据库中的内容
func Fresh() {
	testConn()
}
