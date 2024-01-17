package orm

import (
	"context"
	"log"
	"parsing/orm/ent"
)

// Migrate runs the migration tool
func Migrate() {
	client := NewClient()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Printf("failed closing connection to mysql: %v\n", err)
		}
	}(client)
}
