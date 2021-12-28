go env -w GOBIN=/Users/youdi/go/bin
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct // 使用七牛云的

GOPROXY="goproxy.io,goproxy.cn,direct"

go env // 打印环境设置

go mod init github.com/codeoneline/mygo

go build -ldflags "-X main.gitCommit=09a771a81a3157e4a8f9566110c5635eb23ac879 -X main.gitDate=20211202 -extldflags -Wl,-z,stack-size=0x800000" -trimpath -v -o /home/jsw/go/src/github.com/RichardZHLH/go-ethereum/build/bin/geth ./cmd/geth
go build -trimpath -v -o /home/jsw/go/src/github.com/codeoneline/mygo/build/bin/geth ./cmd/geth
