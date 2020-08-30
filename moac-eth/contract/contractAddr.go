package main

import (
	"fmt"

	"github.com/innowells/moac-lib/common"
	"github.com/innowells/moac-lib/crypto"
)

func main() {
	// 0xc12031318143c148aedea573e1b0f691e33f3092
	account1 := common.HexToAddress("0xc12031318143c148aedea573e1b0f691e33f3092")
	account2 := common.HexToAddress("0xc12031318143c148aedea573e1b0f691e33f3093")
	account3 := common.HexToAddress("0xc12031318143c148aedea573e1b0f691e33f3094")

	for i := 0; i < 3; i++ {
		addr := getCreateAddress(account1, uint64(i))
		fmt.Printf("main--%d--[%s]: %s\n", i, account1.String(), addr.String())
	}
	fmt.Println("")
	for i := 3; i < 8; i++ {
		addr := getCreateAddress(account2, uint64(i))
		fmt.Printf("fund--%d--[%s]: %s\n", i, account2.String(), addr.String())
	}
	fmt.Println("")
	for i := 0; i < 5; i++ {
		addr := getCreateAddress(account3, uint64(i))
		fmt.Printf("fund--%d--[%s]: %s\n", i, account3.String(), addr.String())
	}
}

func getCreateAddress(b common.Address, nonce uint64) common.Address {
	return crypto.CreateAddress(b, nonce)
}
