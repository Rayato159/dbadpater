package database

type Db interface {
	InsertOne()
}

func NewDb(db Db) Db {
	return db
}
