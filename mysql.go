package data_mysql

import (
	"database/sql"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/infrago/data"
	_ "github.com/go-sql-driver/mysql"
)

type (
	mysqlDriver struct{}

	mysqlConnection struct {
		instance *data.Instance
		db       *sql.DB
		actives  int64
	}

	mysqlDialect struct{}
)

func (d *mysqlDriver) Connect(inst *data.Instance) (data.Connection, error) {
	return &mysqlConnection{instance: inst}, nil
}

func (c *mysqlConnection) Open() error {
	dsn := strings.TrimSpace(c.instance.Config.Url)
	if dsn == "" {
		if v, ok := c.instance.Setting["dsn"].(string); ok {
			dsn = v
		}
	}
	if dsn == "" {
		return fmt.Errorf("missing mysql dsn")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return err
	}
	c.db = db
	return nil
}

func (c *mysqlConnection) Close() error {
	if c.db == nil {
		return nil
	}
	err := c.db.Close()
	c.db = nil
	return err
}

func (c *mysqlConnection) Health() data.Health {
	return data.Health{Workload: atomic.LoadInt64(&c.actives)}
}

func (c *mysqlConnection) DB() *sql.DB {
	return c.db
}

func (c *mysqlConnection) Dialect() data.Dialect {
	return mysqlDialect{}
}

func (mysqlDialect) Name() string { return "mysql" }
func (mysqlDialect) Quote(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "`", "")
	return "`" + s + "`"
}
func (mysqlDialect) Placeholder(_ int) string { return "?" }
func (mysqlDialect) SupportsILike() bool      { return false }
func (mysqlDialect) SupportsReturning() bool  { return false }
