package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {

	rootCmd.AddCommand(migrateCmd, crawlerCmd, mapperCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
