package db

import (
// "errors"
)

const (
	MONGO = "mongodb"
	MYSQL = "mysql"
)

// var errtype = errors.New("Database Type not found... ")

// type DBLayer interface {
// 	CountTweets() (int, error)
// 	// GetAuthors() []bson.M
// }

//ConnectDatabase connects to a database type o using the provided connection string
// func ConnectDatabase(o string) (DBLayer, error) {
// 	return NewMongoStore()
// switch o {
// case MONGO:
// 	return NewMongoStore()
// 	// case MYSQL:
// 	// 	return NewMySQLDataStore(cstring)
// }
// log.Println("Could not find ", o)
// return nil, errtype
// }
