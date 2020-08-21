package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	ecies "go.dedis.ch/kyber/encrypt/ecies"
	ed25519 "go.dedis.ch/kyber/group/edwards25519"
	"encoding/hex"
	"fmt"
	"crypto/x509"
)

func main() {
	var privateKey *ecdsa.PrivateKey

	//标准库生成ecdsa密钥对
	privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//
	suite := ed25519.NewBlakeSHA256Ed25519()
	//转为ecies密钥对
	var eciesPrivateKey *ecies.PrivateKey
	var eciesPublicKey *ecies.PublicKey
	eciesPrivateKey = ecies.ImportECDSA(privateKey)
	eciesPublicKey = &eciesPrivateKey.PublicKey

	var message string = "this is a message, 这是需要加密的数据"
	fmt.Println("原始数据: \n" + message)

	//加密
	cipherBytes, _ := ecies.Encrypt(suite, eciesPublicKey, []byte(message), nil, nil)
	//密文编码为16进制字符串输出
	cipherString := hex.EncodeToString(cipherBytes)

	fmt.Println("加密数据: \n" + cipherString)

	//解密
	//decrypeMessageBytes, _ := eciesPrivateKey.Decrypt(suite, cipherBytes, nil, nil)
	bytes, _ := hex.DecodeString(cipherString)
	decrypeMessageBytes, _ := eciesPrivateKey.Decrypt(suite, bytes, nil, nil)
	decrypeMessageString := string(decrypeMessageBytes[:])

	fmt.Println("解密数据: \n" + decrypeMessageString)
}