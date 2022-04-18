module github.com/guessi/mysql-connectivity

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli/v2 v2.4.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220412211240-33da011f77ad
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
)
