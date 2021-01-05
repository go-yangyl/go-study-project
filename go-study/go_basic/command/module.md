##  Go Modules

---

#### go mod 命令

  ```shell script
go mod init 项目名 生成go mod 和 go sum
  ```

#### go mod 文件

  ```go
module example.com/foobar
go 1.13
require (
    example.com/apple v0.1.2
    example.com/banana v1.2.3
    example.com/banana/v2 v2.3.4
    example.com/pineapple v0.0.0-20190924185754-1b0db40df49a
)
exclude example.com/banana v1.2.4
replace example.com/apple v0.1.2 => example.com/rda v0.1.0 
replace example.com/banana => example.com/hugebanana
  ```

- Module : 用于定义当前项目的模块路径

- Go : 用于设置预期的 Go 版本
- require：用于设置一个特定的模块版本
- exclude：用于从使用中排除一个特定的模块版本。
- replace：用于将一个模块版本替换为另外一个模块版本。



#### go sum文件

go.sum 是类似于比如 dep 的 Gopkg.lock 的一类文件，它详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值以备 Go 在今后的操作中保证项目所依赖的那些模块版本不会被篡改。



#### GO111MODULE

- auto：只在项目包含了 go.mod 文件时启用 Go modules，自适应
- on ：必须要使用go mod
- off ：关闭



#### GOPROXY

```htt
go env -w GO111MODULE=on   设置为只能使用go mod
go env -w GOPROXY=https://goproxy.cn,direct  设置go代理
go get url    下载指定的依赖
go get -u     更新现有的依赖
go mod tidy   把go mod的包下载下来
go mod vendor 将go mod里面的包导入到vendor当中
```





