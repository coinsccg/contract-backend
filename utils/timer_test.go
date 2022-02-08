package utils

import (
	"fmt"
	"math/big"
	"testing"
)

func TestTimerBy5Minute(t *testing.T) {
	n := new(big.Int)
	balance, _ := n.SetString("222222222220222222555584558844115285", 10)
	balanceBytesCopy1 := make([]byte, len(balance.Bytes()))
	copy(balanceBytesCopy1, balance.Bytes())
	fmt.Println(balanceBytesCopy1)
	n1 := new(big.Int)
	n1.SetBytes(balanceBytesCopy1)
	fmt.Println(n1)
	n1.Div(n1, big.NewInt(2))
	fmt.Println(n1)
}
