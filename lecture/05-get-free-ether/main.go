package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	url  = "https://goerli.infura.io/v3/5900180607824e62b0f4e67daa646799"
	murl = "https://mainnet.infura.io/v3/5900180607824e62b0f4e67daa646799"
)

func main() {
	//ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	//_, err := ks.NewAccount("password")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = ks.NewAccount("password")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c405d88eef718be968d466b3fadb089c8657bba4
	//d7661531fefeb2cbb78647cf21f8fa048d80ac05
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress("c405d88eef718be968d466b3fadb089c8657bba4")
	a2 := common.HexToAddress("d7661531fefeb2cbb78647cf21f8fa048d80ac05")

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance 1:", b1)
	fmt.Println("Balance 2:", b2)
}
