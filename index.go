package data_mysql

import (
	"github.com/bamgoo/bamgoo"
	"github.com/bamgoo/data"
)

func Driver() data.Driver {
	return &mysqlDriver{}
}

func init() {
	bamgoo.Register("mysql", Driver())
}
