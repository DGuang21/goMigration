package migration

var (
	upMap       = map[string]func(){}
	downMap     = map[string]func(){}
	versionList = []string{}
)

/**
 * @Description:
 * @param up
 * @param down
 * @auth: daguang
 */
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

func Migrate(force ...bool) {
	// diff
	for _, version := range versionList {
		vFunc, ok := upMap[version]
		if ok {
			vFunc()
		}
	}
}

func RollBack(step ...int) {

}

func RollBackAndMigrate(step ...int) {

}

func Fresh() {

}
