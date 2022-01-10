package main

import (
	"encoding/hex"
	"fmt"

	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
)

func main() {
	xc, _ := xuper.New("127.0.0.1:37801", xuper.WithConfigFile("./conf/sdk.yaml"))
	acc, _ := account.GetAccountFromFile("", "")

	args := map[string]string{
		"key": "a",
	}
	tx, err := xc.InvokeWasmContract(acc, "counter", "increase", args)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(tx.Tx.Txid))
}
