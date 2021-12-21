package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"test/StarifySM"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deploy() {
	client, err := ethclient.Dial("https://rpc.testnet.fantom.network/")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("***")
	if err != nil {
		log.Fatal(err)
	}

	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}

	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(4002))
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = big.NewInt(20000000000)
	//auth.Signer =

	input := big.NewInt(10000)
	address, tx, instance, err := StarifySM.DeployStarifySM(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("gas", gasPrice)

	_ = instance
}

func main() {
	deploy()
	//client, err := ethclient.Dial("https://rpc.testnet.fantom.network/")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//address := common.HexToAddress("0x**")
	//instance, err := StarifySM.NewStarifySM(address, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//privateKey, err := crypto.HexToECDSA("****")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}
	//
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(4002))
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)     // in wei
	//auth.GasLimit = uint64(6721975) // in units
	//auth.GasPrice = big.NewInt(20000000000)
	//
	//name, err := instance.Mint(auth, big.NewInt(100))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(name)
}
