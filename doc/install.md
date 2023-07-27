#### install

1. install go
   ```
    wget https://studygolang.com/dl/golang/go1.20.6.linux-amd64.tar.gz
    tar -C /usr/local -zxvf go...
    vi /etc/profile
    # 在/etc/profile最后一行添加
    export GOROOT=/usr/local/go
    export PATH=$PATH:$GOROOT/bin
    # 保存退出后source一下（vi 的使用方法可以自己搜索一下）
    source /etc/profile
   ```
2. install make & protoc
   ```
   wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.5/protoc-3.19.5-linux-x86_64.zip
   (apt install zip)
   unzip protoc...
   拷贝到.../go/bin
   ```
3. make
   make init & make all

#### config

1. install pg
   docker run --name load-book -e POSTGRES_PASSWORD=123456 -p 5432:5432 -d postgres:9.6
2. install redis
   docker run --restart=always --log-opt max-size=100m -p 6379:6379 --name myredis -d redis --appendonly yes

#### run

1. kratos run
2. go run main.go wire_gen.go -conf ../../configs
