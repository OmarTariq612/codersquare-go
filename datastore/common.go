package datastore

import (
	"os"
	"path"

	"github.com/OmarTariq612/codersquare-go/datastore/dao"
	"github.com/OmarTariq612/codersquare-go/datastore/sqlitedb"
)

type Database interface {
	dao.UserDAO
	dao.PostDAO
	dao.LikeDAO
	dao.CommentDAO
	CloseConnection() error
}

// var DB Database = memorydb.NewMemoryDB()
// var DB Database = sqlitedb.NewSqliteDB("/home/omar/programming/go/codersquare-go/datastore/sqlitedb/db/database.sqlite")

var DB Database

func init() {
	wd, err := os.Getwd() // this is the path of the binary
	if err != nil {
		panic(err)
	}
	DB = sqlitedb.NewSqliteDB(path.Join(wd, "datastore", "sqlitedb", "db", "database.sqlite"))
}

func CloseDB() error {
	return DB.CloseConnection()
}
