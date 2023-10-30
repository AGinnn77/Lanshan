### Go环境变量

```
GOROOT    //从Go安装与升级过程可知，这是Go的安装根目录，与JDK的JAVA_HOME一样，包含了官方的依赖，源代码、命令工具和文档等
GOPATH    //类似一个工作空间，在使用GO MOD之前，必须在该目录的src下创建项目，pkg目录存放依赖包，bin目录存放编译的可执行文件；使用GO MOD之后，项目不必强制在该目录下了。
GO111MODULE    //使用启用GO MOD，有三个值：auto、off和on。off和on就不细说，比较明显，一个是不启动，一个是启动。当为auto时，项目中存在go.mod文件或者项目位置在GOPATH之外，则启用GO MOD，否则，则不启用GO MOD。
GOPROXY    //默认的GO依赖下载地址，国内开发者无法正常访问，所以需要设置一个镜像代理，一般可以设置为：https://goproxy.cn,direct
GOOS    //指定Go程序编译时的目标操作系统
GOARCH    //指定Go程序编译时的目标指令集架构
GOENV    //环境变量配置文件
```

参考链接：https://blog.csdn.net/xiaohui249/article/details/132479703#comments_28343489



