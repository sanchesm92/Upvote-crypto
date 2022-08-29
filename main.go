package main

import (
	"upvote-crypto/server"
	"upvote-crypto/database"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}