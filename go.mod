module github.com/bamgoo/data-mysql

go 1.25.3

require (
	github.com/bamgoo/bamgoo v0.0.0
	github.com/bamgoo/data v0.0.0
	github.com/go-sql-driver/mysql v1.8.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/bamgoo/base v0.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
)

replace github.com/bamgoo/bamgoo => ../bamgoo

replace github.com/bamgoo/data => ../data
