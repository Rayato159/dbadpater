package database

import "fmt"

type postgres struct{}

func PostgresConn() Db { return &postgres{} }

func (p *postgres) InsertOne() {
	fmt.Println("INSERT INTO blablabla...")
}
