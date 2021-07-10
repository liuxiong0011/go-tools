package keystore

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

func TestTextHash(t *testing.T) {
	// 02
	//keyfile02 := "D:\\Projects\\src\\github.com\\ethereum\\go-ethereum\\cmd\\geth\\data\\keystore\\UTC--2021-07-04T15-29-13.858301700Z--b982c439d3662d139f4d1563ceb795c6393c15ce"
	// 04
	keyfile := "D:\\Projects\\src\\github.com\\ethereum\\go-ethereum\\cmd\\geth\\data\\keystore\\UTC--2021-07-04T15-37-27.672583400Z--f1ed0d8181489e928f8097b2c4d02102eb189f56"
	password := ".........."
	keyjson, err := ioutil.ReadFile(keyfile)
	if err != nil {
		panic(fmt.Errorf("failed to read the keyfile at '%s': %v", keyfile, err))
	}
	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		panic(fmt.Errorf("error decrypting key: %v", err))
	}
	keystr := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	fmt.Printf(keystr)
}

func TestBip39(t *testing.T) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(fmt.Errorf("failed to NewEntropy %v", err))
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		panic(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0") //最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址
	account, err := wallet.Derive(path, false)
	if err != nil {
		panic(err)
	}

	address := account.Address.Hex()
	privateKey, _ := wallet.PrivateKeyHex(account)
	publicKey, _ := wallet.PublicKeyHex(account)

	fmt.Printf("%s\n", mnemonic)
	fmt.Println("address0:", address)      // id为0的钱包地址
	fmt.Println("privateKey:", privateKey) // 私钥
	fmt.Println("publicKey:", publicKey)   // 公钥

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") //生成id为1的钱包地址
	account, err = wallet.Derive(path, false)
	if err != nil {
		panic(err)
	}

	fmt.Println("address1:", account.Address.Hex())

}
