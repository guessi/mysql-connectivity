FROM public.ecr.aws/docker/library/golang:1.25-alpine3.23 AS BUILDER
RUN apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/guessi/mysql-connectivity
COPY *.go go.mod go.sum ./
RUN GOPROXY=direct go build -o /go/bin/mysql-connectivity

FROM public.ecr.aws/docker/library/alpine:3.23
COPY --from=BUILDER /go/bin/mysql-connectivity /opt/
CMD ["/opt/mysql-connectivity"]
