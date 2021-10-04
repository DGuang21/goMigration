package migration

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

// CleanMigration
// Clean up funcs that are permanently up/down in memory
// 清理在初始化时常驻在内存中的 up/down 方法
func CleanMigration() {
	upList = []func(){}
	downList = []func(){}
}

func Migrate(force ...bool) {

}

func RollBack(step ...int) {

}

func RollBackAndMigrate(step ...int) {

}

func Fresh() {

}