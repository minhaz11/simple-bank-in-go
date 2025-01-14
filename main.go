package main

import "log"

func main() {
	dbStore, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
		return
	}

	if err = dbStore.Seeder(); err != nil {
		log.Fatal(err)
		return
	}

	server := NewApiServer(":8085", dbStore)

	server.Run()
}