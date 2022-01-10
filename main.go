package main

import (
	"encoding/hex"
	"fmt"

	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
)

func main() {
	xc, _ := xuper.New("127.0.0.1:37801")
	acc, _ := account.GetAccountFromFile("", "")

	t, err := xc.Transfer(acc, "", "1")
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(t.Tx.GetTxid()))
}
