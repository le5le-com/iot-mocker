# le5le/iot-mocker

le5le/iot-mocker 是乐吾乐物联网平台数据模拟工具。

# 配置

配置优先级：  
./my.yaml > ./config.yaml

# 运行

```
./iot-mocker

# TLS 证书

自己生成，在 cert 目录下面执行：

```

openssl genrsa -out server.key 2048

openssl req -new -x509 -key server.key -out server.crt -days 3650

```

生产环境需要机构认可

# exe 签名

防止病毒误杀。

需要下载[Windows-SDK](https://developer.microsoft.com/en-us/windows/downloads/windows-sdk/)，设置 PATH 环境变量：C:\Program Files (x86)\Windows Kits\10\bin\10.0.22621.0\x64\

```

// 证书生成一次就行，不用重复生成
// 生成证书，需要设置密码
makecert -n "CN=le5le.com,E=admin@le5le.com,C=China" -r app.cer -ss le5le.com -sv app.pvk -$ individual

// 转换证书格式 1
cert2spc app.cer app.spc
// 转换证书格式 2
pvk2pfx -pvk app.pvk -spc app.spc -pfx app.pfx

// 管理员身份命令行执行
signtool sign /fd SHA256 /f app.pfx iot-mocker.exe
// 带时间戳签名
signtool sign /fd SHA256 /f app.pfx /t http://timestamp.digicert.com iot-mocker.exe

```

```
