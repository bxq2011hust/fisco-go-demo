package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"credit-letter/contract"
)

const (
	issuerPEM   = "./accounts/0x83309d045a19c44dc3722d15a6abd472f95866ac.pem"
	holderPEM   = "./accounts/0xb112a7f71afc39229623b694f266ed2e97fe9d1d.pem"
	acceptorPEM = "./accounts/0xb112a7f71afc39229623b694f266ed2e97fe9d1d.pem"
)

func parseECPrivateKeyFromPEM(file string) (*ecdsa.PrivateKey, error) {
	keyHex, _, _, err := conf.LoadECPrivateKeyFromPEM(file)
	if err != nil {
		fmt.Printf("parse private key failed, err: %v\n", err)
		return nil, err
	}
	return crypto.HexToECDSA(keyHex)
}

func main() {
	// 解析私钥，私钥用于交易签名
	keyHex, _, _, err := conf.LoadECPrivateKeyFromPEM(issuerPEM)
	if err != nil {
		fmt.Printf("parse private key failed, err: %v\n", err)
		return
	}
	// 构造config对象
	config := &conf.Config{IsHTTP: false, ChainID: 1, CAFile: "ca.crt", Key: "sdk.key", Cert: "sdk.crt",
		IsSMCrypto: false, GroupID: 1, PrivateKey: keyHex, NodeURL: "127.0.0.1:20200"}
	// 创建client
	client, err := client.Dial(config)
	if err != nil {
		fmt.Printf("Dial Client failed, err:%v", err)
		return
	}
	// 部署CreditLetterFactory合约
	address, _, instance, err := contract.DeployCreditLetterFactory(client.GetTransactOpts(), client)
	if err != nil {
		fmt.Printf("Deploy failed, err:%v", err)
		return
	}
	// 部署合约得到的地址应该存储
	fmt.Println("CreditLetterFactory address: ", address.Hex())

	credirLetterFactory := &contract.CreditLetterFactorySession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	issuer, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		fmt.Printf("parse private key failed, err: %v\n", err)
		return
	}
	issuerAddress := crypto.PubkeyToAddress(issuer.PublicKey)
	holder, err := parseECPrivateKeyFromPEM(holderPEM)
	if err != nil {
		fmt.Printf("parse private key failed, err: %v\n", err)
		return
	}
	holderAddress := crypto.PubkeyToAddress(holder.PublicKey)

	acceptor, err := parseECPrivateKeyFromPEM(acceptorPEM)
	if err != nil {
		fmt.Printf("parse private key failed, err: %v\n", err)
		return
	}
	acceptorAddress := crypto.PubkeyToAddress(acceptor.PublicKey)
	issuanceTime := big.NewInt(time.Now().Unix())
	interestRate := big.NewInt(1)
	credit := big.NewInt(10000000)
	_, receipt, err := credirLetterFactory.Create(issuerAddress, holderAddress, acceptorAddress, issuanceTime, interestRate, credit)
	if err != nil {
		fmt.Printf("Create credirLetter failed, err: %v\n", err)
		return
	}
	// 解析abi
	creditLetterFactoryABI, err := abi.JSON(strings.NewReader(contract.CreditLetterFactoryABI))
	if err != nil {
		fmt.Printf("parse abi failed, err: %v\n", err)
		return
	}
	// 使用creditLetterFactoryABI 解析返回值：新创建信用证的地址
	var creditLetterAddress common.Address
	err = creditLetterFactoryABI.Unpack(&creditLetterAddress, "create", common.FromHex(receipt.Output))
	if err != nil {
		fmt.Printf("parse return value failed, err: %v\n", err)
		return
	}
	fmt.Println("new credit letter address:", creditLetterAddress.Hex())
	// 使用地址创建CreditLetter实例
	creditLetterInstance, err := contract.NewCreditLetter(creditLetterAddress, client)
	if err != nil {
		fmt.Printf("NewCreditLetter failed, err: %v\n", err)
		return
	}

	creditLetter := &contract.CreditLetterSession{Contract: creditLetterInstance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	// 调用GetInfo()
	ret, err := creditLetter.GetInfo()
	if err != nil {
		fmt.Printf("creditLetter.GetInfo() failed, err: %v\n", err)
		return
	}
	fmt.Printf("creditLetter.GetInfo() result:\n %+v\n", ret)
}
