package main

import (
	"context"
	"fmt"
	"log"
	"twitter/config"
	"twitter/postgres"
)

func main() {
	fmt.Println("Hello World")

	ctx := context.Background()
	conf := config.NewConfig()
	log.Println("config=", conf)
	db := postgres.New(ctx, conf)
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected to database")
	log.Println("working")
}

//14
