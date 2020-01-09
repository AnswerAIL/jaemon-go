# Go

> 1. go language environment installation	

````bash
	# 安装 GO 语言
    sudo apt-get install golang -y
    
    # 查看 go 版本信息
    go version
    
    # 查看go的环境变量
    go env
    
    # 查看go的安装目录
    go env | grep "GOROOT"
    
    
    # Ubuntu升级GO版本
     sudo apt-get remove golang-1.6
     sudo apt autoremove
     wget https://www.golangtc.com/static/go/1.9.2/go1.9.2.linux-amd64.tar.gz
     tar -C /usr/local -zxvf go1.9.2.linux-amd64.tar.gz
     echo "export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin" >> ~/.profile
     source ~/.profile
     go version
    
    
    sudo apt install golang-go 
    sudo apt-get install golang-1.9 
    sudo apt-get remove golang-1.9 
    sudo apt autoremove 

    # 先卸载旧版本 
    curl -O https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz 
    tar -C /usr/local -zxvf go1.9.linux-amd64.tar.gz 
    mkdir -p ~/go/src 
    echo "export GOPATH=$HOME/go" >> ~/.profile 
    echo "export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin" >> ~/.profile 
    source ~/.bashrc 
    source ~/.profile 
    go version 
````



> 2. the first go program 

````go
package main

import "fmt"

func main() {
  var a int = 10
  var b = 20
  c := 30
  var str string = "hello world."
  fmt.Println(a, b, c)
  fmt.Println(str)
}

/*
	1: vim hello.go 
	2: go run hello.go
*/
````



> 3. 函数和方法   [Refer]([https://blog.csdn.net/wo198711203217/article/details/57402613)

````go
package main

import "fmt"

type AnswerFunc struct {}

// 函数
func main() {
  var af AnswerFunc
  var rlt  int
  rlt =  af.max(100, 23);
  fmt.Printf("max: %d\n", rlt)

  m, n := af.swap("11", "22")
  fmt.Printf("m=%s, n=%s\n", m, n)
}

// 方法-带有接收器 (t *AnswerFunc)
func (t *AnswerFunc) max(x int, y int) int {
  t.aprint("answer")    // success
 // aprint("answer")    // failed
  if (x > y) {
    return x
  } else {
    return y
  }
}

func (t *AnswerFunc) swap(x, y string) (string, string) {
  return y, x
}

func (t *AnswerFunc) aprint(msg string) {
  fmt.Println("hello: " + msg)
}
````



> ### 4. 循环

````go
package main
import "fmt"
func main() {

  nums := []int{2, 3, 4}
  sum := 0
  // 字符'_'省略了序号
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)
  
  newSum := 0
  // i为遍历的序号下标, 从0开始
  for i, num := range nums {
    newSum += num
    fmt.Println(i)
  }
  fmt.Println("newSum:", newSum)
}


/** ##########################################################*/

package main

import "fmt"

func main() {
    a := [...]int{1, 2, 3, 5, 6, 7, 8, 9}

     for i := 0; i < len(a); i=i+2  {
        fmt.Println(a[i])
        fmt.Println(a[i+1])
        fmt.Println("----------------")
    }
   
}
````



> ###  5. 数字和字符串互转

````go
package main
import ("strconv")

func main() {
    // 数字字符串转数字
    i, err := strconv.Atoi("12345")
    if err != nil {
        panic(err)
    }
    i += 3
    println(i)
   // 数字转字符串
    s := strconv.Itoa(12345)
    s += "3"
    println(s)
}
````



> ### 6. 指针

````go
package main

import "fmt"

func main() {
   var a int= 20      /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

  /* 指针变量的存储地址 */
   ip = &a 

   fmt.Printf("a 变量的地址是: %x\n", &a)

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量的存储地址: %x\n", ip)

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip)
   
   // 当一个指针被定义后没有分配到任何变量时, 它的值为 nil, nil 指针也称为空指针 
   var  ptr *int
   fmt.Printf("ptr 的值为 : %x\n", ptr)
   if (ptr == nil){
     fmt.Println("ptr is nil pointer.")
   }
    
   var  pstr *string
   fmt.Printf("pstr 的值为 : %x\n", pstr)
}

/*
	a 变量的地址是: c82000a288
    ip 变量的存储地址: c82000a288
    *ip 变量的值: 20
    ptr 的值为 : 0
    ptr is nil pointer.
    pstr 的值为 : 0
*/
/*
	Refer: https://studygolang.com/articles/5878
	指针变量前加'*'用来获取指针所指向的内容
	变量前加'&'用来获取变量的内存地址
*/ 
````



> ### 7. go test

````go
// mymath.go
package mymath

func Add(x, y int) int {
    return x + y
}

func Minus(x, y int) int {
    return x - y
}



// mymath_test.go
package mymath_test

import "mymath"
import "testing"


func TestAdd(t *testing.T) {
    ret := mymath.Add(2, 3)
    if ret != 5 {
        t.Error("Expected 5, got ", ret)
    }
}

func TestMinus(t *testing.T) {
    ret := mymath.Minus(2, 3)
    if ret != -1 {
        t.Error("Expected -1, got ", ret)
    }
}

/*
	execute: go test
	output: 
		PASS
		ok  	mymath	0.002s
	
	Attention: 文件在mymath目录下, mymath目录在 $GOPATH 的src目录下
*/
````



> ### 8. Map

````go
package main

import(
    "fmt"
)

func main() {
  // 创建map对象方式1, paramMap := make(map[string]bool){}
  paramMap := map[string]bool{
    "xiwei": true,
    "vfuchong": false,
    "kayak": false,
    "simope": false,
  }

  for name, flag  := range paramMap {
    fmt.Printf("%s \t %t\n", name, flag)
  }

  if _, ok := paramMap["xiwei"]; ok {
    fmt.Println("xiwei is ok.", )
  }
  
  // 创建map对象方式2
  info := make(map[string]string)
  info["name"] = "answer"
  info["base"] = "sz"
  for key, value := range info {
    fmt.Println("key: " + key + ", value: " + value)
  }
    
}

/* Output: 
	kayak 	 false
    simope 	 false
    xiwei 	 true
    vfuchong 	 false
    xiwei is ok.
    key: name, value: answer
	key: base, value: sz
*/
````



> ### 9. 切片

````go
package main

import "fmt"

func main() {
  s := make([]string, 10, 12)
  // len 长度超过了 cap, cap 以2的指数型扩容
  fmt.Printf("s len: %d, cap: %d, slice=%v\n", len(s), cap(s), s)
  // 向切片 s 追加元素
  s = append(s, "100", "200", "300")
  fmt.Printf("s len: %d, cap: %d, slice=%v\n", len(s), cap(s), s)

  p := []string{"1", "2"}
  fmt.Printf("p len: %d, cap: %d, slice=%v\n", len(p), cap(p), p)
 
  // 一个切片在未初始化之前默认为 nil, 长度为 0, 即空切片
  var q []string
  fmt.Printf("q len: %d, cap: %d, slice=%v\n", len(q), cap(q), q)

  numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Println("numbers = ", numbers)
  // 索引从 0 开始, 前闭后开 
  fmt.Println("numbers[1: 4] = ", numbers[1: 4])
  fmt.Println("numbers[:3] = ", numbers[:3])
  fmt.Println("numbers[3:] = ", numbers[3:])
}

/*	Output: 
	s len: 10, cap: 12, slice=[         ]
    s len: 13, cap: 24, slice=[          100 200 300]
    p len: 2, cap: 2, slice=[1 2]
    q len: 0, cap: 0, slice=[]
    numbers =  [1 2 3 4 5 6 7 8 9]
    numbers[1: 4] =  [2 3 4]
    numbers[:3] =  [1 2 3]
    numbers[3:] =  [4 5 6 7 8 9]
*/
````



> ### 10. defer

````go
package main

import "fmt"

func main() {
  defer_test()
}

/*
  defer通常用来释放函数内部变量
  是在函数正常返回之后执行
*/
func defer_test() {
  fmt.Println("defer test start...")
  i := 0
  defer fmt.Println("defer 1: ", i)
  i++
  defer fmt.Println("defer 2: ", i)
  i++
  defer fmt.Println("defer 3: ", i)
  i++
  defer fmt.Println("defer 4: ", i)
  i++
  defer fmt.Println("defer 5: ", i)
  fmt.Println("defer test end!!!")
  return
}

/* Output: 
	defer test start...
    defer test end!!!
    defer 5:  4
    defer 4:  3
    defer 3:  2
    defer 2:  1
    defer 1:  0

*/
````



### JSON 与 MAP 互转

```go
package main

/* go run json_to_map.go */

import ( 
   "encoding/json"
   "fmt"
)

func main() { 
   jsonStr := "{\"address\":\"北京\",\"username\":\"answer\"}"
   myMap := make(map[string]string) 
   // JSON字符串转map
   json.Unmarshal([]byte(jsonStr),&myMap) 
   fmt.Println(myMap) 
   fmt.Println(myMap["address"])
   fmt.Println(myMap["username"])

   myMap["address"] = "深圳"
   fmt.Println(myMap["address"])

   // map 转 JSON 字符串
   jsonBytes, err := json.Marshal(myMap)
   if err != nil {
      fmt.Println(err)
    }
   fmt.Println(string(jsonBytes))
}





package main

/* go run json_to_map.go */

import ( 
   "encoding/json"
   "fmt"
)

func main() { 
   modifyJsonStr := "{\"address\":\"sz\",\"username\":\"answerCoder\",\"size\":\"10000\"}"
   queryJsonStr := "{\"address\":\"北京\",\"username\":\"answer\"}"

   queryMap := make(map[string]string) 
   // JSON字符串转map
   json.Unmarshal([]byte(queryJsonStr),&queryMap) 

   fmt.Println(queryMap) 
   fmt.Println(queryMap["address"])
   fmt.Println(queryMap["username"])

   queryMap["address"] = "深圳"
   fmt.Println(queryMap["address"])

   modifyMap := make(map[string]string) 
   // JSON字符串转map
   json.Unmarshal([]byte(modifyJsonStr),&modifyMap)
   // 变量 modifyMap 中的信息
   for key, value  := range modifyMap {
   //  fmt.Printf("%s \t %s\n", key, value)
     queryMap[key] = value 
  }

   // map 转 JSON 字符串
   jsonBytes, err := json.Marshal(queryMap)
   if err != nil {
      fmt.Println(err)
    }
   fmt.Println(string(jsonBytes))
}
```





### JSON 和结构体互转

```go
package main 

import ( 
  "fmt" 
  "encoding/json" 
) 

type appInfo struct { 
  Appid string `json:"appId"` 
} 

type response struct { 
  RespCode string `json:"respCode"` 
  RespMsg string `json:"respMsg"` 
  AppInfo appInfo `json:"app"` 
} 

type JsonResult struct { 
  Resp response `json:"resp"` 
} 

func main() { 
  jsonstr := `{"resp": {"respCode": "000000","respMsg": "成功","app": {"appId": "d12abd3da59d47e6bf13893ec43730b8"}}}`

  var JsonRes JsonResult 
  json.Unmarshal([]byte(jsonstr), &JsonRes) 
  fmt.Println("After Parse: ", JsonRes) 
  fmt.Println("RespCode: ", JsonRes.Resp.RespCode)
  fmt.Println("RespMsg: ", JsonRes.Resp.RespMsg)
  fmt.Println("Appid: ", JsonRes.Resp.AppInfo.Appid)

  fmt.Println("\n\n\n###########################\n\n\n")

  transJsonStr, _ := json.Marshal(JsonRes)
  fmt.Println(string(transJsonStr))
}



/**
	运行： go run json_to_struts.go
	输出：
		After Parse:  {{000000 成功 {d12abd3da59d47e6bf13893ec43730b8}}}
		RespCode:  000000
		RespMsg:  成功
		Appid:  d12abd3da59d47e6bf13893ec43730b8



		###########################



		{"resp":{"respCode":"000000","respMsg":"成功","app":		{"appId":"d12abd3da59d47e6bf13893ec43730b8"}}}


*/
```



### 解析JSON报文

```go
package main 
  {
        "header": {
                "orgNo": "A000001", 
                "operator": "jane", 
                "transDate": "20181119", 
                "transTime": "105101"
        }, 
        "body": {
                "transCode": "Create_Order", 
                "orderNos": [
                        "201811191051010000001",
                        "201811191051010000002",
                        "201811191051010000003"
                ], 
                "respCode": "0000", 
                "respDesc": "success"
        }
  }
*/

import ( 
  "fmt" 
  "encoding/json" 
) 

type XWResponse struct {
  Header HeaderTag `json:"header"`
  Body BodyTag `json:"body"`
}

type HeaderTag struct {
  OrgNo string `json:"orgNo"`
  Operator string `json:"operator"`
  TransDate string `json:"transDate"`
  TransTime string `json:"transTime"`
}

type BodyTag struct {
  TransCode string `json:"transCode"`
  OrderNos []string `json:"orderNos"`
  RespCode string `json:"respCode"`
  RespDesc string `json:"respDesc"`
}



func main() { 
  respStr := `{ "header": { "orgNo": "A000001", "operator": "jane", "transDate": "20181119", "transTime": "105101" }, "body": { "transCode": "Create_Order", "orderNos": ["201811191051010000001", "201811191051010000002", "201811191051010000003"], "respCode": "0000", "respDesc": "success"} }`

  var response XWResponse 
  json.Unmarshal([]byte(respStr), &response) 
  fmt.Println("After Parse: ", response) 
  fmt.Println("orgNo: ", response.Header.OrgNo)
  fmt.Println("orders: ", response.Body.OrderNos) 

  var orderNos []string 
  orderNos=response.Body.OrderNos
  for index, orderNo := range orderNos {
    fmt.Printf("index: %+v, orderNo: %+v.\n", index, orderNo)
  } 

  fmt.Println("\n\n\n###########################\n\n\n")

  transJsonStr, _ := json.Marshal(response)
  fmt.Println(string(transJsonStr))
}



/**
	运行： go run json2struts.go
	输出：
    After Parse:  {{A000001 jane 20181119 105101} {Create_Order [201811191051010000001 201811191051010000002 201811191051010000003] 0000 success}}
    orgNo:  A000001
    orders:  [201811191051010000001 201811191051010000002 201811191051010000003]
    index: 0, orderNo: 201811191051010000001.
    index: 1, orderNo: 201811191051010000002.
    index: 2, orderNo: 201811191051010000003.



###########################



	{"header":	{"orgNo":"A000001","operator":"jane","transDate":"20181119","transTime":"105101"},"body":{"transCode":"Create_Order","orderNos":["201811191051010000001","201811191051010000002","201811191051010000003"],"respCode":"0000","respDesc":"success"}}





*/
```











### go 命令行操作

- [x] 编译软件包

  ````bash
  go build hello.go	
  ````
  - `-x`：可以打印出执行过程的详细信息，辅助调试 
  - `-gcflags` :  指定编译器参数 
  - `-ldflags`： 指定链接器参数，常见的可以通过 -X 来动态指定包变量值 

- [x] 清理项目, 删除编译生成的二进制文件和临时文件 

  ```bash
  go clean
  ```

  - `-i`：删除 go install 安装的文件 
  - `-n` ：打印删除命令，而不执行，方便进行测试检查 
  - `-r`：递归清除，对依赖包也执行清理工作 
  - `-x`：执行清除过程同时打印执行的删除命令，方便进行测试检查 

- [x] 对代码进行格式化检查和修正 

  ````bash
  gofmt hello.go
  go fmt hello.go
  ````

  - `-d`：仅显示不符合格式规定的地方，不进行修正 
  - `-e`：打印完整错误内容，默认是只打印 10 行 
  - `-l`：列出不符合格式规定的文件路径 
  - `-r`：重写的规则 
  - `-s`：对代码尝试进行简化 
  - `-w`：对不符合默认风格的代码进行修正 

- [x] 对代码进行格式风格检查，打印出不符合 Go 语言推荐风格的代码 

  ````bash
  # 安装 golint
  apt-get install golint
  # 检查当前目录下的文件代码风格
  golint .
  # 检查指定文件
  golint  hello.go
  
  # 对超级账本 Fabric 所有代码进行风格检查, 注意后面的 ... 代表递归检查所有子目录下内容
  golint $GOPATH/src/github.com/hyperledger/fabric/...
  ````

- [ ] 获取导入包列表

  ````go
  go list -f "{{ join .Imports "\n"}}" 
  ````

- [x] [其他命令参考地址](https://yeasy.gitbooks.io/blockchain_guide/content/appendix/golang/tools.html)





### Go编写规范

```bash
# 公共函数: 函数名首字母大写
# 私有函数: 函数名首字母小写
```





### Golang 编辑器中导入GitHub上的库

```bash
# 配置好 GO 的环境变量 GOROOT 和 GOPATH

# 进入cmd， 执行 go get github.com/hyperledger/fabric/core/chaincode/shim
# 回车之后， 会自动下载项目到GOPATH中的src目录下
# 或者
# 在IDE工具中的对应库 Alt+Enter自动导入对应的项目

############################################
# 			切换库的版本
############################################
# 进入到对应的库目录下， 即： GOPATH src目录下的库
# 执行 git branch -r
# git checkout -b release-1.3 origin/release-1.3
```





































































