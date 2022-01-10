package main

import "github.com/xuperchain/xuper-sdk-go/v2/account"

func getAliceAccount() *account.Account {
	acc, err := account.GetAccountFromFile("", "")
	if err != nil {
		panic(err)
	}

	return acc
}
