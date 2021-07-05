package keystore

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestTextHash(t *testing.T) {
	// 02
	//keyfile02 := "D:\\Projects\\src\\github.com\\ethereum\\go-ethereum\\cmd\\geth\\data\\keystore\\UTC--2021-07-04T15-29-13.858301700Z--b982c439d3662d139f4d1563ceb795c6393c15ce"
	// 04
	keyfile := "D:\\Projects\\src\\github.com\\ethereum\\go-ethereum\\cmd\\geth\\data\\keystore\\UTC--2021-07-04T15-37-27.672583400Z--f1ed0d8181489e928f8097b2c4d02102eb189f56"
	password := ".........."
	keyjson, err := ioutil.ReadFile(keyfile)
	if err != nil {
		panic(fmt.Errorf("failed to read the keyfile at '%s': %v", keyfile, err))1
	}
	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		panic(fmt.Errorf("error decrypting key: %v", err))
	}
	keystr := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	fmt.Printf(keystr)
}
