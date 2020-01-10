# GO 语言

___

&nbsp;

# 使用 logrus 日志框架
```bash
# 获取代码, 在 $GOPATH/src/github.com目录下
go get github.com/sirupsen/logrus
# eg. C:\Users\answer\go\src\github.com\sirupsen\logrus


```

&nbsp;

# proto文件生成go代码
```bash
# 新增对应的 ***.proto 文件， 执行以下命令
protoc.exe  --go_out=./  common.ptoto

# 生成对应的 ***.ptoto.pb.go 文件
```

&nbsp;

# int16, int32, int64等类型说明
 - int16： 16位整数(16bit integer)，相当于short, 占2个字节, -32768 ~ 32767
 - int32： 32位整数(32bit integer), 相当于 int, 占4个字节, -2147483648 ~ 2147483647
 - int64： 64位整数(64bit interger), 相当于 long long, 占8个字节, -9223372036854775808 ~ 9223372036854775807
 - byte：  相当于byte(unsigned char), 0 ~ 255




&nbsp;

# Reference
 - [windows 下 goprotobuf 的安装与使用](https://blog.csdn.net/u010979642/article/details/103896533)
 - [Go 语言入门教程](http://c.biancheng.net/golang/)