FROM golang:1.18.1-alpine


WORKDIR /build
ENV GOPROXY https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
RUN go mod init http-server
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o http-server .
EXPOSE 80

ENTRYPOINT ["/build/http-server"]