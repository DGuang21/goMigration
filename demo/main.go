package main

import (
	"fmt"
	"github.com/DGuang21/goMigration/migration"
	"time"
)

func main() {
	t1 := time.Now()
	Migrate()
	fmt.Println(time.Now().Sub(t1))
	migration.CleanMigration()
}
