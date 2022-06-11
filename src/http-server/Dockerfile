FROM golang:1.18.1-alpine as build

WORKDIR /httpserver/
ENV GOPROXY https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
RUN go mod init http-server
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o http-server .

FROM busybox
EXPOSE 80
COPY --from=build  /httpserver /httpserver
WORKDIR /httpserver
ENTRYPOINT ["./http-server"]