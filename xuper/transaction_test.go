package xuper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuperchain/service/pb"
)

func TestA(t *testing.T) {
	ccc := string([]byte{128})
	// fmt.Println(string([]byte{127}))
	// ccc := "abc你好{}【】，。、？ßßßå≈∂é"
	fmt.Println("AAAA:", len(ccc))
	fmt.Println("AAAA:", []byte(ccc))
	args := map[string]string{"a": ccc}
	input, err := json.Marshal(args)
	if err != nil {
		panic("!")
	}
	fmt.Println("QQQ:", input)
	fmt.Println("QQQ:", string(input))

	fmt.Println(args["a"])

	var a = map[string]string{}
	err = json.Unmarshal(input, &a)
	if err != nil {
		panic(err)
	}
	c := a["a"]
	// c := []byte(b.(string))
	fmt.Println("BBB", len([]byte(c)))
	fmt.Println("BBB", []byte(c))
	fmt.Println("BBB", []byte(string([]byte(c))))
	fmt.Println(c)
}

func TestTransaction(t *testing.T) {
	// acc, _ := account.CreateAccount(1, 1)
	acc1, _ := account.CreateAccount(1, 1)

	type Case struct {
		hasHash       bool
		signAcc       *account.Account
		hasErr        bool
		inAuthRequire bool
	}

	cases := []*Case{
		{
			hasHash:       true,
			signAcc:       nil,
			hasErr:        true,
			inAuthRequire: true,
		},
		{
			hasHash:       false,
			signAcc:       acc1,
			hasErr:        false,
			inAuthRequire: true,
		},
		{
			hasHash:       true,
			signAcc:       acc1,
			hasErr:        true,
			inAuthRequire: false,
		},
	}

	for _, c := range cases {
		tx := &Transaction{
			Tx: &pb.Transaction{},
		}
		if c.hasHash {
			tx.DigestHash = []byte("haha")
		}

		if c.inAuthRequire {
			tx.Tx.AuthRequire = append(tx.Tx.AuthRequire, acc1.GetAuthRequire())
		}

		err := tx.Sign(c.signAcc)
		if c.hasErr {
			if err == nil {
				t.Error("Transactions assert failed1")
			}
		} else {
			if err != nil {
				t.Error("Transactions assert failed2")
			}
		}
	}

}

func TestComplianceCheck(t *testing.T) {

}
