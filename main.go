package main

import (
	_ "github.com/go-sql-driver/mysql"
	"parsing/orm"
	"parsing/parse"
)

func main() {
	//orm.Migrate()

	c := orm.NewClient()
	defer c.Close()

	s := orm.NewStore(c)
	p := parse.NewCrawler(s)

	p.Run()
}
