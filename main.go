package main

import "github.com/Rayato159/clean-db-implementation/database"

func main() {
	db1 := database.Db(database.MongoConn())
	db1.InsertOne()

	db2 := database.Db(database.PostgresConn())
	db2.InsertOne()
}
