### GOPATH 的弊端

- 没有版本控制概念
- 无法同步一致第三方版本号
- 无法指定当前项目引用的第三方版本号

### Go Modules 模式

- go mod 命令
    - go mod init 生成 go.mod 文件
    - go mod download 下载 go.mod 文件中指明的所有依赖
    - go mod tidy 整理现有的依赖
    - go mod graph 查看现有的依赖结构
    - go mod edit 编辑 go.mod 文件
    - go mod vendor 导出项目所有的依赖到 vendor 目录
    - go mod verify 校验一个模块是否被篡改过
    - go mod why 查看为什么需要依赖模块

- go mod 相关环境变量
    - GO111MODULE 是否开启 go modules 模式，建议 go v1.11 之后都设置为 on
    - GOPROXY 项目的第三方依赖库下载源地址，direct 用于 Go 回源到模块版本的源地址去抓去（比如 github 等），该项在国内无法访问的情况下可以设置为国内镜像源
    - GOSUMDB 用于校验拉去的第三方库是否完整，默认也是国外的网站，如果设置了 GOPROXY，这个可以不用设置了
    - GONOPROXY 不需要代理下载的地址
    - GONOSUMDB 不需要校验的下载地址
    - GOPRIVATE 私有的源仓库（不会走代理和校验），go env -w GOPRIVATE="git.example.com,github.com/lustormstout/xxx"
      表示这两个地址是私有仓库，不需要走代理去下载也不需要校验，最后三个一般只设置 GOPRIVATE 就可以了，支持通配符配置

- 查看环境变量 go env
- 设置一项环境变量：go env -w GO111MODULE=on
        
