package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"path/filepath"

	"github.com/innowells/moac-lib/common"
	"github.com/innowells/moac-lib/crypto"
)

type extarg struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

var (
	big0 = big.NewInt(0)
)

func main() {
	key := "00000000000000000000000000000000000000000000000000000000000001aa"
	getAccount(key)
	// getAccountAndCheckBalance(key)
}

func getAccount(key string) {
	seckey := common.Hex2Bytes(key)
	//fmt.Printf("seckey=%s\n", common.Bytes2Hex(seckey))
	var pub ecdsa.PublicKey
	c := crypto.S256()
	pub.X, pub.Y = c.ScalarBaseMult(seckey)
	pub.Curve = c
	addr := crypto.PubkeyToAddress(pub)
	fmt.Printf("%s\n", addr.String())
}

func getAccountAndCheckBalance(key string) {
	baseDir, _ := filepath.Abs("./")
	filepath := filepath.Join(baseDir, "Info")
	//fmt.Printf("filepath = %s\n", filepath)
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//defer f.Close()
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	}

	for i := 0; i < 1; i++ {
		seckey := common.Hex2Bytes(key)
		fmt.Printf("seckey=%s\n", common.Bytes2Hex(seckey))
		var pub ecdsa.PublicKey
		c := crypto.S256()
		pub.X, pub.Y = c.ScalarBaseMult(seckey)
		pub.Curve = c
		addr := crypto.PubkeyToAddress(pub)
		fmt.Printf("%s\n", addr.String())
		addrString := addr.String()
		url := "https://api.etherscan.io/api?module=account&action=balance&address=" + addrString + "&tag=latest&apikey=YourApiKeyToken"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("http.Get err: " + err.Error())
			return
		}
		ext := new(extarg)
		body, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
		err = json.Unmarshal(body, ext)
		if err != nil {
			fmt.Println("json.Unmarshal err: " + err.Error())
			return
		}

		fmt.Printf("seckey = %s\n", common.Bytes2Hex(seckey))
		fmt.Printf("seckey = %s\n", addrString+" : "+ext.Result)
		if ext != nil && len(ext.Result) > 1 {
			appendToFile(f, common.Bytes2Hex(seckey))
			appendToFile(f, addrString+" : "+ext.Result)
			// } else if i%1000 == 0 {
			// 	fmt.Println(strconv.FormatUint(i, 10))
			// 	appendToFile(f, common.Bytes2Hex(seckey))
		}
	}
}

func appendToFile(f *os.File, content string) error {
	// 以只写的模式，打开文件
	// 查找文件末尾的偏移量
	//n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	_, err := f.WriteString(content + "\n")
	if err != nil {
		fmt.Println("WriteString. err: " + err.Error())
	}
	return err
}
