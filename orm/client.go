package orm

import (
	"log"
	"parsing/orm/ent"
)

func NewClient() *ent.Client {
	client, err := ent.Open("mysql", "root:nopass@tcp(localhost:3306)/crawler?parseTime=True")
	if err != nil {
		log.Printf("failed opening connection to mysql: %v\n", err)
	}

	return client
}
