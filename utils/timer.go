package utils

import (
	"auction/config"
	"auction/constant"
	"auction/token"
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	"time"

	"auction/logs"
	"auction/models"
)

// Timer 定时器
func Timer() {
	// 记录第一次时间
	if err := models.UpdateRefreshTime(time.Now().Add(time.Hour * 4).Unix()); err != nil {
		logs.GetLogger().Error(err)
		return
	}
	TimerBy7Minute()
	// 定时任务1 每7分钟拉取全网持币地址
	ticker1 := time.NewTicker(time.Minute * 5)
	go func() {
		for {
			<-ticker1.C
			TimerBy7Minute()
		}
	}()

	// 定时任务2 每4小时发放奖励
	ticker2 := time.NewTicker(time.Hour * 4)
	go func() {
		for {
			<-ticker2.C
			TimerBy4Hours()
		}
	}()
}

func TimerBy7Minute() {
	// 发送请求
	var (
		page   = 1
		offset = 200
	)
	for {
		//var tmp = map[string]interface{}{
		//	"status":  "1",
		//	"message": "ok",
		//	"result": []map[string]interface{}{
		//		{
		//			"TokenHolderAddress":  "0xf977814e90da44bfa03b6295a0616a897441acec",
		//			"TokenHolderQuantity": "10000000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0xf977814e90da44bfa03b6295a0616a897441acec",
		//			"TokenHolderQuantity": "20000000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0xf977814e90da44bfa03b6295a0616a897441acec",
		//			"TokenHolderQuantity": "3000000000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0xb1b55c88798bd5cbfd7dc144ed3ba55f1cb61746",
		//			"TokenHolderQuantity": "1000210000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x04b35d8eb17729b2c4a4224d07727e2f71283b73",
		//			"TokenHolderQuantity": "10002100000000000000001",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x5f62c1a5950492d3d59b70c94eb344c923125629",
		//			"TokenHolderQuantity": "10002100000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0xf31d585c18e411b956388cad051f31be235a854e",
		//			"TokenHolderQuantity": "10002100000000004120000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x2012dd3e0491d03b07ca06b173efef80bf7caf36",
		//			"TokenHolderQuantity": "10002100000002541000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x418aa6bf98a2b2bc93779f810330d88cde488888",
		//			"TokenHolderQuantity": "100021000120000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x66b870ddf78c975af5cd8edc6de25eca81791de1",
		//			"TokenHolderQuantity": "1000214000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0xb980255fb2f3e68a4fa9d8575e32199b2d9858db",
		//			"TokenHolderQuantity": "100210000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x2c789cf18dd172dad075e5af94f59b84bc16895d",
		//			"TokenHolderQuantity": "1010210000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x37a59ec9241877fba10fae1831f9b448342e6ffc",
		//			"TokenHolderQuantity": "1000210000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0xd15a9a3a1b833552d57b5e11e963bdf1ecabe084",
		//			"TokenHolderQuantity": "25210000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x4fa5f07604a7b58d3f9d668ec4ca4af9459893d7",
		//			"TokenHolderQuantity": "1660210000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0xbd7170971505828f17a48f0177f9203da7b57d14",
		//			"TokenHolderQuantity": "1777210000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x20b3df3b0dca3b72cbc5f210c920a87bf4132775",
		//			"TokenHolderQuantity": "100041000005000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x4865c4c96d0ac0f912108572a4c514ea94738e6e",
		//			"TokenHolderQuantity": "5410000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0xcba2fcfd833ddb870c52b9db26ba888bec8a2faa",
		//			"TokenHolderQuantity": "23510000000000000000000",
		//		},
		//		{
		//			"TokenHolderAddress":  "0x63e8d12d21bad8e8fb94b220b8b55113e29ddb9a",
		//			"TokenHolderQuantity": "54870000000000000000000",
		//		},
		//	},
		//}
		//
		//bytes, err := json.Marshal(tmp)

		bytes, err := RequestBSCHoldersListByContractAddress(constant.CONTRACT_ADDRESS, constant.BSC_API_KEY, page, offset)
		if err != nil {
			logs.GetLogger().Error(err)
			break
		}
		var resp *models.Resp
		if err = json.Unmarshal(bytes, &resp); err != nil {
			logs.GetLogger().Error(err)
			break
		}

		if err = models.InsertHolder(resp); err != nil {
			logs.GetLogger().Error(err)
			break
		}

		if len(resp.Result) < offset {
			break
		}
		page++
	}
}

func TimerBy4Hours() {
	holders, err := models.FindCurrentHoldersRank()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	// 连接rpc
	client, err := ethclient.Dial(constant.BSC_RPC_POD)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	// 获取私钥、账户、json string
	_, _, key := token.Keystore(config.GetConfig().PrivateKey)
	eclient := token.NewEthClient(client)

	// 查询竞拍池余额
	balance, err := eclient.ERC20BalanceOf(constant.CONTRACT_ADDRESS, constant.AUCTION_ADDRESS)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	if balance.Cmp(big.NewInt(10000)) == -1 {
		if err = models.UpdateRefreshTime(time.Now().Add(time.Hour * 4).Unix()); err != nil {
			logs.GetLogger().Error(err)
			return
		}
		return
	}

	//n := new(big.Int)
	//balance, _ := n.SetString("50004154721525142247122", 10)

	// 前20名随机发放奖励
	rand.Seed(time.Now().Unix())

	var num = 100
	m := len(holders)
	if m > 20 {
		m = 20
	}

	balanceBytesCopy := make([]byte, len(balance.Bytes()))
	copy(balanceBytesCopy, balance.Bytes())

	for i, v := range holders {
		v.LastReward = "0"
		if i < 3 {
			n1 := new(big.Int)
			n1.SetBytes(balanceBytesCopy)
			n1.Mul(n1, big.NewInt(3-int64(i)))
			n1.Div(n1, big.NewInt(10))
			if err = eclient.ERC20Transfer(constant.CONTRACT_ADDRESS, key, v.HolderAddress, n1); err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			v.LastReward = n1.String()
		}

		random := rand.Intn(num/(m-i)-1) + 1
		if i == m-1 {
			random = num
		}
		n2 := new(big.Int)
		n2.SetBytes(balanceBytesCopy)
		n2.Mul(n2, big.NewInt(4*int64(random)))
		n2.Div(n2, big.NewInt(1000))
		if err = eclient.ERC20Transfer(constant.CONTRACT_ADDRESS, key, v.HolderAddress, n2); err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		n3 := new(big.Int)
		lastReward, _ := n3.SetString(v.LastReward, 10)
		n3.Add(lastReward, n2)
		v.LastReward = n3.String()
		num -= random
		if i >= 19 {
			break
		}
	}

	if err = models.UpdateLastReward(holders); err != nil {
		logs.GetLogger().Error(err)
		return
	}
	if err = models.UpdateRefreshTime(time.Now().Add(time.Hour * 4).Unix()); err != nil {
		logs.GetLogger().Error(err)
		return
	}
}
