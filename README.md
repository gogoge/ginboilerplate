# 用Golang/Gin 實作API Server

## 工作計劃

- [x] 本機安裝Golang
- [x] IDE設定
    - [x] VS Code
- [x] 基本語法略懂
- [x] go tool基本指令
    - [x] go get
    - [x] go build
    - [x] go install
- [x] 套件管理
    - [x] govendor
- [ ] gin
    - [x] 基本server
    - [x] server file structure
        - [x] import / export
    - [ ] DB操作
    - [ ] JWT
    - [ ] Cookie
- [ ] 參數
    - [ ] go flag
    - [ ] viper
- [x] Security
    - [x] secure http headers
- [ ] 測試
    - [ ] Unit test
    - [ ] API test
- [ ] Dockerize

## 前置準備  

### Golang語法  
https://tour.golang.org/list  
https://golang.org/doc/code.html#Organization  

### 安裝

[MacOS安裝](doc/installation/macos.md)  


## 套件管理
Golang的所有程式，都會統一放在一個WorkSpace裡，  
環境變數$GOPATH必需指向這個WorkSpace  

WorkSpace會有三個資料夾`bin`, `pkg`與`src`  
透過`go get`下載套件時，會同時將套件的source code放置在src中  

如果希望像Node.js一樣，Project管理各自的套件  
要將套件存放在各Project下的`vendor`資料夾  

```
/$GOPATH
   |-----/bin
   |-----/pkg
   |-----/src
   |      |----/Project1
   |      |       |-------/vendor   
   |      |       |-------main.go
```

在node.js中我們使用npm作套件管理，在golang中，還沒有像npm一樣，官方的工具  
這個案子中選用的是`govendor`，各工具比較可以參閱 [go依赖包管理工具对比](https://ieevee.com/tech/2017/07/10/go-import.html)

### 安裝Govendor

```
go get -u github.com/kardianos/govendor
```

### govendor的使用方式

1. 在Project執行init

> 類似 npm init或yarn init  

```
govendor init
```
init後會產生vendor資料夾跟vendor.json檔

2. 下載指定套件，並記錄到vendor.json

> 類似 npm i 或 yarn add  

```
govendor fetch <套件URL>
```

3. 以vendor.json下載套件

> 類似 npm 或 yarn

```
govendor sync
```

### 設定gitignore

```
vendor/*/
```

### Golang尋找套件的次序

1. Project/vendor
2. 向上層目錄找，直到找到src/vendor
3. $GOPATH
4. $GOROOT

## Gin Server

```
govendor fetch github.com/gin-gonic/gin
```

新增main.go

```
package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```

執行

```
go run main.go
```

再以browser連到localhost:8080/ping，若有拿到pong代表server建立成功

### 檔案結構

參考下面這篇文章
https://github.com/josephspurrier/gowebapp

在Project Root有一個entry point `main.go`  
其他都放在app目錄中，以import的方式匯入  

https://peter.bourgon.org/go-best-practices-2016/
上文的Best Practice中提到

> Always use fully-qualified import paths. Never use relative imports.

為了讓程式中可以import子資料夾下的go檔，寫的時候必需以完整路徑，不建議用`./`  
有些Project會把source code放到vendor下，但vendor的用途類似node_modules，應該只放3rd prarty的lib  
不建議這樣放


```
import app/shared/secure  // 不好
import devpack.cc/ginboilerplate/app/shared/secure  //好
```


資料夾結構如下  
```
/$GOPATH
   |-----/bin
   |-----/pkg
   |-----/src
   |      |----/Project1
   |      |       |-----/app
   |      |       |       |-----/controller
   |      |       |       |-----/route
   |      |       |       |-----/model
   |      |       |       |-----/shared
   |      |       |                |-----/secure
   |      |       |-------/vendor   
   |      |       |-------main.go
```
vendor下除了govendor fetch下來的packages外，app目錄下是所有的source codea

## secure http header
使用的是 https://github.com/unrolled/secure  
請參照 [vendor/app/shared/secure/secure.go](vendor/app/shared/secure/secure.go)