# go-quick-start

## 安装 go 环境

最简单的方式是，直接去官网下载安装包: 

安装包下载地址为：https://golang.org/dl/。

如果打不开可以使用这个地址：https://golang.google.cn/dl/。

查看go命令路径:

```sh
$ which go
/usr/local/go/bin/go
```

查看 go 语言安装版本: 

```sh
$ go version  
go version go1.12.7 darwin/amd64
```

可以看到，目前安装的是 go1.12.7 版本。


## 第一个 HelloWord 程序

创建一个目录名为 go-quick-start, 在目录中创建一个 main.go 文件: 

```sh
$ mkdir go-quick-start
$ touch main.go
$ vi main.go
```

代码文件 main.go 内容如下: 

```java
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!");
}
```

## 编译并运行代码

(1) 直接使用 go run 命令运行 main.go 程序:

```sh
$ go run main.go 
Hello, World!
```

(2) 使用 go build 构建代码，并执行编译后的可执行文件:

```sh
$ go build main.go
$ ./main 
Hello, World!
```


## 创建 go modules 模块: 

大多数编程语言都会有包管理工具，像Node有`npm`，PHP有`composer`，Java有`Maven`和`Gradle`。 
可是，Go语言一直缺乏一个官方的包管理(`Dep`被称为官方试验品`official experiment`)。
终于，在`go1.11` 版本中，新增了`module`管理模块功能，用来管理依赖包。

```sh
$ go mod init github.com/cnwyt/go-quick-start
go: creating new go.mod: module github.com/cnwyt/go-quick-start
```

会生成一个 go.mod 文件，指明了模块名和 go 的版本: 

```sh
module github.com/cnwyt/go-quick-start
go 1.12
```

## 引入一个第三方包 uniplaces/carbon: 

可以使用 go get 直接下载包: 

```sh
$ go get github.com/uniplaces/carbon
```

或者使用 go mod edit 命令编辑mod文件，使用 go mod download 来下载包: 

```sh
$ go mod edit -require github.com/uniplaces/carbon@latest
$ go mod download
```

注意使用 go mod edit 必须指定版本 `path@version`，如 `xxx@latest` 或 `xxx@v0.1.6`， 否则会提示如下错误: 

```sh
$ go mod edit -require github.com/uniplaces/carbon
$ go mod: -require=github.com/uniplaces/carbon: need path@version
```

查看 go.mod 文件: 

```sh
module github.com/cnwyt/go-quick-start

go 1.12

require github.com/uniplaces/carbon v0.1.6
```

可以看到 多了一行 require 语句，指定第三方包的url路径和版本号。

在 main.go 中使用 import 来引入包: 

```java
package main

import "fmt"
import "github.com/uniplaces/carbon"

func main() {
    fmt.Println("Hello, World!");
    fmt.Printf("Unix timestamp:  %d \n", time.Now().Unix())
    fmt.Printf("Right now is:  %s \n", carbon.Now().DateTimeString())
}
```

使用 go run 命令编译并执行 main.go 文件: 

```sh
$ go run main.go 
Hello, World!
Unix timestamp:  1562766341 
当前时间是:  2019-07-10 21:45:41
```

## 使用第三方包操作 MySQL 数据库 

再引入一个第三方包 go-sql-driver/mysql:

```sh
$ go mod edit -require github.com/go-sql-driver/mysql@latest
$ go mod download  
```

查看 go.mod 文件，require 部分已经指定了 go-sql-driver/mysql 包，多个包默认是用圆括号包起来的。

```java
module github.com/cnwyt/go-quick-start

go 1.12

require (
	github.com/go-sql-driver/mysql latest
	github.com/uniplaces/carbon v0.1.6
)
```

修改 main.go 代码:

```java
package main

import (
	"fmt"
	"time"
	"log"
	"github.com/uniplaces/carbon"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Hello, World!");

    fmt.Printf("Unix timestamp:  %d \n", time.Now().Unix())
    fmt.Printf("Right now is:  %s \n", carbon.Now().DateTimeString())

    // 连接数据库
	db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/test")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(3 * time.Second)
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Fatal("----> 数据库连接失败.")
	    panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("----> 数据库连接成功.");
	}
	defer db.Close() 
}
```

运行代码，结果如下: 

```sh
$ go run main.go
Hello, World!
Unix timestamp:  1562766341 
当前时间是:  2019-07-10 21:45:41
----> 数据库连接成功.
```


[END]