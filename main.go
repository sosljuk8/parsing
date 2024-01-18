package main

import (
	_ "github.com/go-sql-driver/mysql"
	"parsing/orm"
	"parsing/parse"
	"parsing/parse/plugin"
)

func main() {
	//orm.Migrate()

	c := orm.NewClient()
	defer c.Close()

	s := orm.NewStore(c)
	jobs := plugin.NewJobs()
	jobs.Add(plugin.NewHydacJob())
	p := parse.NewCrawler(s, jobs)

	p.Run()
}
