package database

import "fmt"

type mongo struct{}

func MongoConn() Db { return &mongo{} }

func (m *mongo) InsertOne() {
	fmt.Println("db.insertOne() blablabla...")
}
