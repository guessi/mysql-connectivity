FROM public.ecr.aws/docker/library/golang:1.17-alpine3.15 as BUILDER
RUN apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/guessi/mysql-connectivity
COPY . .
RUN go build -o /go/bin/mysql-connectivity

FROM public.ecr.aws/docker/library/alpine:3.15
COPY --from=BUILDER /go/bin/mysql-connectivity /opt/
CMD ["/opt/mysql-connectivity"]
