# 使用说明

## 拷贝sdk与节点建立连接所需的证书私钥

拷贝sdk.key、sdk.crt、ca.crt到本目录下

## 编译

```bash
go build main.go
```

## 运行

```bash
./main
```

正确配置，会有如下输出

```bash
CreditLetterFactory address:  0x32CA0eEe658Acd8d92FF68EdAf991B818b6870b2
new credit letter address: 0x7107C7F38fA27e933640893a305B0d6062FCb62E
creditLetter.GetInfo() result:
 {IssuerVal:[131 48 157 4 90 25 196 77 195 114 45 21 166 171 212 114 249 88 102 172] HolderVal:[177 18 167 247 26 252 57 34 150 35 182 148 242 102 237 46 151 254 157 29] AcceptorVal:[177 18 167 247 26 252 57 34 150 35 182 148 242 102 237 46 151 254 157 29] IssuanceTimeVal:+1599721540 InterestRateVal:+1 CreditVal:+10000000 Approved:false Paid:false}
```

## 说明

1. contract目录下的go文件是使用abigen工具生成的，请参考[这里](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/sdk/go_sdk/contractExamples.html)
2. 合约需要使用solc编译得到abi和bin文件，编译得到的这些文件用于abigen生成合约的go代码
3. 使用国密模式，需要使用对应的sm2p256v1的私钥，使用配套的solc的国密版本生成对应的bin和abi，请[参考这里](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/sdk/go_sdk/contractExamples.html#id14)