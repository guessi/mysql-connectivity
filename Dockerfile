FROM public.ecr.aws/docker/library/golang:1.18-alpine3.16 as BUILDER
RUN apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/guessi/mysql-connectivity
COPY *.go go.mod go.sum ./
RUN go build -o /go/bin/mysql-connectivity

FROM public.ecr.aws/docker/library/alpine:3.16
COPY --from=BUILDER /go/bin/mysql-connectivity /opt/
CMD ["/opt/mysql-connectivity"]
