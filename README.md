# googleAuth
googleAuthenticator 谷歌身份验证器，即谷歌动态口令，Google身份验证器Google Authenticator是谷歌推出的基于时间的一次性密码(Time-based One-time Password，简称TOTP)，只需要在手机上安装该APP，就可以生成一个随着时间变化的一次性密码，用于帐户验证。

# 原理
## 秘钥生成

> 1、时间戳，精确到微秒，除以1000，除以30（动态6位数字每30秒变化一次） 
>
> 2、对时间戳余数 hmac_sha1 编码 
>
> 3、然后 base32 encode 标准编码 
>
> 4、输出大写字符串，即秘钥 
>


### 伪代码
```
 func GoogleAuthenticatorCode(string secret)      
 {
    key := base32decode(secret)   
    message := floor(current Unix time / 30)      
    hash := HMAC-SHA1(key, message)      
    offset := last nibble of hash      
    truncatedHash := hash[offset..offset+3]  
    // 4 bytes starting at the offset      
    // Set the first bit of truncatedHash to zero  
    // remove the most significant bit      
    code := truncatedHash mod 1000000      
    // pad code with 0 until length of code is 6      
    return code
 }
```

## 动态码生成
> 1.基于密钥和时间计算一个HMAC-SHA1的hash值，这个hash是160 bit的，
>
> 2.将这个hash值随机取连续的4个字节生成32位整数，
>
> 3.将整数取31位
>
> 4.取模得到一个的整数。

### 伪代码

```
  func GoogleAuthenticatorCode(string secret){
   key := base32decode(secret)      
   message := counter encoded on 8 bytes      
   hash := HMAC-SHA1(key, message)      
   offset := last nibble of hash      
   truncatedHash := hash[offset..offset+3]  
   // 4 bytes starting at the offset      
   // Set the first bit of truncatedHash to zero  
   // remove the most significant bit      
   code := truncatedHash mod 1000000      
   // pad code with 0 until length of code is 6      
   return code
  }
```

## 使用
```
go get -u https://github.com/crazykun/googleAuth
```


## 示例

```
package main

import (
	"fmt"

	"github.com/crazykun/googleAuth"
)

func main() {

	str := `TESTTESTTESTTEST`
	code, _ := googleAuth.GetCode(str)

	fmt.Println(code) 

}
```

JumpServer使用ssh脚本实现一键登录, 自动输入google二次验证
```
go build example/main.go -o googleAuth



#!/bin/sh
code=`./googleAuth TESTTESTTEST`

export LC_CTYPE="en_US.UTF-8"
expect -c "
spawn ssh root@127.0.0.2 -p22
expect \"*ode]:\"
send \"$code\r\"
expect \"*Opt> \"
send \"1\r\"
expect \"*Host]>\"
send \"server_name\r\"
expect \"*]$\"
interact
"
```



