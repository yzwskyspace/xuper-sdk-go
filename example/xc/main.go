package main

import (
	"fmt"

	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
)

const mnemonic1 = "协 盆 折 毅 脱 齿 戴 逻 牛 织 手 门"

func main() {
	a, _ := account.RetrieveAccount(mnemonic1, 1)
	fmt.Println(a.Address)
	testInitiatorSign()
	account.GetAccountFromPlainFile("./alice/alice/")
}

func testInitiatorSign() {
	// // 节点地址。
	node := "10.12.199.82:8101"

	// // 创建节点客户端。
	xclient, _ := xuper.New(node)

	from, err := account.RetrieveAccount("玉 脸 驱 协 介 跨 尔 籍 杆 伏 愈 即", 1)
	if err != nil {
		fmt.Printf("retrieveAccount err: %v\n", err)
		return
	}
	fmt.Println(from.Address)
	tx, err := xclient.Transfer(from, "aaa", "1")
	if err != nil {
		panic(err)
	}

	fmt.Println(tx.Tx)
}
