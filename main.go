package main

import (
	_ "github.com/go-sql-driver/mysql"
	"parsing/orm"
)

func main() {
	orm.Migrate()
}
