package data_mysql

import (
	"github.com/infrago/infra"
	"github.com/infrago/data"
)

func Driver() data.Driver {
	return &mysqlDriver{}
}

func init() {
	infra.Register("mysql", Driver())
}
