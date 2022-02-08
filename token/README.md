###golang调用智能合约
go-ethereum项目文档https://geth.ethereum.org/docs/dapp/native-bindings


####1.安装gcc
```
# 安装gcc
https://jmeubank.github.io/tdm-gcc/
# 添加环境变量
C:\TDM-GCC-64\bin
```

####2.安装工具
```bash
# 克隆项目
git clone https://github.com/ethereum/go-ethereum.git
# 编译
go build ./cmd/abigen
```

####3.通过合约abi生成go文件
```bash
abigen --abi token.abi --pkg main --type Token --out token.go
```