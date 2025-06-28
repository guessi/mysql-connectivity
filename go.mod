module github.com/guessi/mysql-connectivity

go 1.24.0

toolchain go1.24.4

require (
	github.com/go-sql-driver/mysql v1.9.2
	github.com/sirupsen/logrus v1.9.3
	github.com/urfave/cli/v3 v3.3.3
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
