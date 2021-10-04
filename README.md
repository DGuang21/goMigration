# goMigration
基于 Golang 的数据库迁移工具,目前仍在开发中,有兴趣的小伙伴可以联系我一起~

### 食用方法
```
go get https://github.com/DGuang21/goMigration
```
手动将其安装

可通过 ``` gom gen create_c_user_table ``` 方法生成如下文件,在闭包函数中自由修改你所需的字段及方法
```
type migrationCreateUserTableInterface struct {}

func init() {
	export := migrationCreateUserTableInterface{}
	migration.RegisterMigration(export.Up,export.Down)
}

func (m *migrationCreateUserTableInterface) Up() {
	migration.Create("asd","20211004202802_c_up_date", func(table *migration.MigrationTable) {
		table.String("1").NullTable(false).Comment("s")
		table.Timestamps().Comment()
		table.Done()
	})
}

func (m *migrationCreateUserTableInterface) Down() {
	migration.DropIfExists("kos")
}
```
而后执行``` gom run migration ``` 或 ``` go run demo/main.go ``` 进行数据库或表迁移