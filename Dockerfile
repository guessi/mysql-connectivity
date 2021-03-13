FROM golang:1.16-alpine3.13 as BUILDER
RUN apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/guessi/mysql-connectivity
COPY . .
RUN go build -o /go/bin/mysql-connectivity

FROM alpine:3.13
COPY --from=BUILDER /go/bin/mysql-connectivity /opt/
CMD ["/opt/mysql-connectivity"]
