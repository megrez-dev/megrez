# builder
FROM golang:1.19 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o megrez main.go

ADD build/megrez-linux-amd64 /usr/bin/

# runner
FROM alpine

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/megrez /usr/local/bin/megrez

ENTRYPOINT ["/usr/local/bin/megrez"]