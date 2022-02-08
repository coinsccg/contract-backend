package token

import (
	"auction/logs"
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	client *ethclient.Client
}

func NewEthClient(client *ethclient.Client) *EthClient {
	return &EthClient{client}
}

// Keystore
func Keystore(privateKey string) (*ecdsa.PrivateKey, accounts.Account, string) {
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	// 生成新账号
	//pk, _ := crypto.GenerateKey()
	//pkBytes := crypto.FromECDSA(pk)
	//privateKey := hexutil.Encode(pkBytes)
	//fmt.Println(privateKey)
	//account, _ := ks.ImportECDSA(pk, "123")
	//content, _ := ioutil.ReadFile(account.URL.Path)
	//key := string(content)

	// 从密钥导入账号
	//privateKey := ""
	pk, _ := crypto.HexToECDSA(privateKey)
	account, _ := ks.ImportECDSA(pk, "123") // 密码可以随意设置 只有第一次加载密钥才有URL.PATH
	content, _ := ioutil.ReadFile(account.URL.Path)
	os.RemoveAll(account.URL.Path)
	key := string(content)
	return pk, account, key
}

// transfer ETH转账
func (e *EthClient) Transfer(account accounts.Account, pk *ecdsa.PrivateKey, value float64) {
	nonce, err := e.client.PendingNonceAt(context.Background(), account.Address)
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	to := common.HexToAddress("0xe21b202A99Db4AA2f25D7b7a8012928c213ee225")
	transaction := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      200000,
		To:       &to,
		Value:    EthToWei(value),
	}
	tx := types.NewTx(&transaction)
	chainId, err := e.client.NetworkID(context.Background())
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), pk)
	err = e.client.SendTransaction(context.Background(), signTx)
	fmt.Println(err)
}

// EthToWei eth转wei
func EthToWei(value float64) *big.Int {
	val := big.NewFloat(value)
	base := new(big.Float)
	base.SetInt(big.NewInt(1 << 18))
	val.Mul(val, base)
	result := new(big.Int)
	val.Int(result)
	return result
}

// ERC20Transfer ERC20转账
func (e *EthClient) ERC20Transfer(contractAddress, key, to string, num *big.Int) error {
	contract, err := NewToken(common.HexToAddress(contractAddress), e.client)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	// 根据keystore签署交易事务
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123", big.NewInt(97))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	// 调用合约转账
	_, err = contract.Transfer(auth, common.HexToAddress(to), num)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

// ERC20BalanceOf 查询ERC20余额
func (e *EthClient) ERC20BalanceOf(contractAddress, to string) (*big.Int, error) {
	contract, err := NewToken(common.HexToAddress(contractAddress), e.client)
	if err != nil {
		logs.GetLogger().Error(err)
		return big.NewInt(0), err
	}
	balance, err := contract.BalanceOf(nil, common.HexToAddress(to))
	if err != nil {
		logs.GetLogger().Error(err)
		return big.NewInt(0), err
	}
	return balance, nil
}
