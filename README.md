## Megrez 博客

> Megrez 为北斗之一的天权星，古称文曲星，作为博客项目的名字再适合不过了

## 快速开始

### Windows

click https://github.com/AlkaidChan/megrez/releases/download/0.1.0-alpha.1/megrez-windows-amd64.exe

```bash
$ ./megrez-windows-amd64.exe
```


### MacOS

```bash
$ wget --no-check-certificate https://github.com/AlkaidChan/megrez/releases/download/0.1.0-alpha.1/megrez-darwin-amd64
$ chmod +x megrez-darwin-amd64
$ ./megrez-darwin-amd64
```

### Linux

```bash
$ wget --no-check-certificate https://github.com/AlkaidChan/megrez/releases/download/0.1.0-alpha.1/megrez-linux-amd64
$ chmod +x megrez-linux-amd64
$ ./megrez-linux-amd64
```

### Docker

```bash
$ docker run -it -d --name megrez -p 8080:8080 alkaidchen/megrez
```

## 编译运行
> Golang: 1.19.3

```bash
$ git clone https://github.com/AlkaidChan/megrez.git
$ git submodule init
$ git submodule update
$ go mod tidy
$ go run main.go
```