# MySQL Connectivity Check Utility

[![GoDoc](https://godoc.org/github.com/guessi/mysql-connectivity?status.svg)](https://godoc.org/github.com/guessi/mysql-connectivity)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/mysql-connectivity)](https://goreportcard.com/report/github.com/guessi/mysql-connectivity)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/mysql-connectivity)](https://github.com/guessi/mysql-connectivity/blob/master/go.mod)

[![Docker Stars](https://img.shields.io/docker/stars/guessi/mysql-connectivity.svg)](https://hub.docker.com/r/guessi/mysql-connectivity/)
[![Docker Pulls](https://img.shields.io/docker/pulls/guessi/mysql-connectivity.svg)](https://hub.docker.com/r/guessi/mysql-connectivity/)
[![Build Status](https://cloud.drone.io/api/badges/guessi/mysql-connectivity/status.svg)](https://cloud.drone.io/guessi/mysql-connectivity)

# Use Case

To evaluate downtime during MySQL cluster maintenance, we can use this simple utility to measure the downtime

# Usage

Check MySQL connectivity with command "SELECT 1"

```
./mysql-connectivity \
    --user tester \
    --pass mySecr3tP@ssword \
    --addr mysql.example.dev \
    --dbname testdb \
    --command "SELECT 1"
```

Sample Output:

```
INFO[2021-03-13T23:55:45+08:00] Opening database connection
INFO[2021-03-13T23:55:45+08:00] Ping pass
DEBU[2021-03-13T23:55:46+08:00] 1
DEBU[2021-03-13T23:55:47+08:00] 1
DEBU[2021-03-13T23:55:48+08:00] 1
DEBU[2021-03-13T23:55:49+08:00] 1
DEBU[2021-03-13T23:55:50+08:00] 1
DEBU[2021-03-13T23:55:51+08:00] 1
```

# How to get developer build?

    $ go get -u github.com/guessi/mysql-connectivity
    $ go install github.com/guessi/mysql-connectivity

# License

[Apache-2.0](LICENSE)
