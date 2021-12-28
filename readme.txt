go env -w GOBIN=/Users/youdi/go/bin
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct // 使用七牛云的

GOPROXY="goproxy.io,goproxy.cn,direct"

go env // 打印环境设置

go mod init github.com/codeoneline/mygo
