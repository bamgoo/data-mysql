# data-mysql

`data-mysql` 是 `data` 模块的 `mysql` 驱动。

## 安装

```bash
go get github.com/infrago/data@latest
go get github.com/infrago/data-mysql@latest
```

## 接入

```go
import (
    _ "github.com/infrago/data"
    _ "github.com/infrago/data-mysql"
    "github.com/infrago/infra"
)

func main() {
    infra.Run()
}
```

## 配置示例

```toml
[data]
driver = "mysql"
```

## 公开 API（摘自源码）

- `func Driver() data.Driver`
- `func (d *mysqlDriver) Connect(inst *data.Instance) (data.Connection, error)`
- `func (c *mysqlConnection) Open() error`
- `func (c *mysqlConnection) Close() error`
- `func (c *mysqlConnection) Health() data.Health`
- `func (c *mysqlConnection) DB() *sql.DB`
- `func (c *mysqlConnection) Dialect() data.Dialect`
- `func (mysqlDialect) Name() string { return "mysql" }`
- `func (mysqlDialect) Quote(s string) string`
- `func (mysqlDialect) Placeholder(_ int) string { return "?" }`
- `func (mysqlDialect) SupportsILike() bool      { return false }`
- `func (mysqlDialect) SupportsReturning() bool  { return false }`

## 排错

- driver 未生效：确认模块段 `driver` 值与驱动名一致
- 连接失败：检查 endpoint/host/port/鉴权配置
