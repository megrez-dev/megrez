# builder
FROM golang:1.19 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o megrez main.go

# runner
FROM alpine

ENV TZ Asia/Shanghai

RUN apk add tzdata && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && apk del tzdata

WORKDIR /app

COPY --from=builder /app/megrez /usr/local/bin/megrez

ENTRYPOINT ["/usr/local/bin/megrez"]