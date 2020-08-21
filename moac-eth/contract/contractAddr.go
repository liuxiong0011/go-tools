package main

import (
	"fmt"

	"github.com/innowells/moac-lib/common"
	"github.com/innowells/moac-lib/crypto"
)

func main() {
	account := common.HexToAddress("0xcB8775C1726c797a1855E86b9F22A7fa0CB7532D") //0xc12031318143c148aedea573e1b0f691e33f3092
	addr := getCreateAddress(account, 1)
	fmt.Printf("ContractAddress: %s\n", addr.String())
}

func getCreateAddress(b common.Address, nonce uint64) common.Address {
	return crypto.CreateAddress(b, nonce)
}
