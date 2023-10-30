## Go语言学习

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



### Go的基本语法
#### main包
```go
package main	//每个执行的go程序，都需要有一个main包，main包里需要有一个main函数
```
#### 声明变量
```go
//声明变量
var a int	//整数类型的默认值为0
var b = 10
var c int = 20	//声明并赋值
var (
	d string
	e int
)

func main(){
    y := "abc"	//定义局部变量简便方式：声明+赋值
}
```

#### 声明常量

```go
const a = 10	//常量声明的同时必须赋值
const(
	b = 20
    c		//未赋值变量的值与上一个变量一样
    d
)

const(
	m = 1
	n = 2
	x = iota	//自带的特殊常量：计算器（从0开始计数），计算是第几个常量
	y
)
```

#### 基本输入输出

```go
//fmt：标准输入输出包
//bufio：也可以接受输入输出

fmt.Scan()		//接受用户的输入，用空格分隔分别传给变量，敲完所需的输入才会结束
fmt.Scan(&a, &b)

fmt.Scanf()		//严格按照指定格式输入，否则报错
fmt.Scanf("a=%d,b=%s", &a, &b)

fmt.Scanln()	//与fmt.Scan()一样用空格分隔分别传给变量，但以换行符来结束本次输入
fmt.Scanln(&a, &b)

fmt.Print()		//原样输出，不换行
{
	fmt.Print("hello")
	fmt.Print("world")		//"helloworld"
}

fmt.Println()	//自动添加换行符和空格分隔符
fmt.Println("hello","world")	//"hello world"

fmt.Printf()	//格式化输出
{
    name := "sctl"
    fmt.Printf("hello %s", name)	//"hello sctl"
}
```

参考链接：https://blog.csdn.net/m0_69298614/article/details/129899970

#### 函数定义

```go
//典型定义
func add(x int, y int) int {
	return x + y
}

//函数可以接受0个或多个参数，需要指定类型
func greet (name string) {
    fmt.Println("Hello, " + name)
}

//函数可以返回一个或多个值，需要指定类型；如果没有返回值，可以将返回类型留空
func addAndMultiply(x, y int) (int, int) {
    sum := x+y
    product :=x * y
    return sum,product
}
```

#### 基础数据结构等

参考链接：https://blog.csdn.net/lxw1844912514/article/details/117705689

类型不匹配：https://blog.csdn.net/qq_44838906/article/details/108310341

#### 利用并发和通道实现在函数调用后输入参数

在 Go 语言中，函数调用时的参数是在函数调用之前确定的，因此无法在函数调用后再输入参数。这是由于函数调用是按照代码的执行顺序进行的。

如果你希望在函数调用后输入参数，可以考虑使用并发（concurrency）和通道（channel）来实现。通过创建一个 goroutine 来处理函数调用，并在另一个 goroutine 中接收输入参数。

下面是一个使用并发和通道的示例代码：

```
package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)

    go func() {
        fmt.Print("请输入一个整数：")
        var num int
        _, err := fmt.Scan(&num)
        if err != nil {
            fmt.Println("输入错误：", err)
            ch <- 0 // 发送一个零值到通道，表示错误
            return
        }
        ch <- num // 发送输入的整数到通道
    }()

    num := <-ch // 从通道接收输入的整数
    if num != 0 {
        result := double(num)
        fmt.Println("结果：", result)
    }
}

func double(n int) int {
    return n * 2
}
```

在上面的代码中，我们创建了一个通道 `ch`，用于在两个 goroutine 之间传递数据。在匿名函数中启动了一个 goroutine，用于接收键盘输入，并将输入的整数发送到通道 `ch` 中。在主 goroutine 中，我们使用 `<-ch` 语法从通道中接收输入的整数。如果接收到的值不等于零，则调用 `double` 函数进行处理并打印结果。

通过使用并发和通道，我们可以在函数调用后等待输入参数，以实现在函数调用后再输入参数的效果。不过需要注意的是，这种方式增加了代码的复杂性，对并发和通道的理解要求较高。因此，在实际开发中，需要根据具体需求和场景来选择合适的解决方案。





### math/rand包的使用

参考链接：https://blog.csdn.net/boysoft2002/article/details/126337102

```go
rand.Int()	//返回一个随机整数
rand.Intn(n)	//返回[0,n)间的非负整数
```

