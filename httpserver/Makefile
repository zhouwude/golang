export tag=v2.0
root:
	export ROOT=github.com/cncamp/golang

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .
# 依赖 build -t 自定义镜像 name 和 tag
release: build
	echo "building httpserver container"
	docker build -t zhouwude/httpserver:${tag} . 
# 
push: release #这里没有仓库地址默认推送到 docker.io
	echo "pushing zhouwude/httpserver"
	docker push zhouwude/httpserver:v2.0
