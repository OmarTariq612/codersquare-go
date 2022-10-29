package datastore

import (
	"github.com/OmarTariq612/codersquare-go/datastore/dao"
	"github.com/OmarTariq612/codersquare-go/datastore/sqlitedb"
)

type Database interface {
	dao.UserDAO
	dao.PostDAO
	dao.LikeDAO
	dao.CommentDAO
}

// var DB Database = memorydb.NewMemoryDB()

var DB Database = sqlitedb.NewSqliteDB("/home/omar/programming/go/codersquare-go/datastore/sqlitedb/db/database.sqlite")
