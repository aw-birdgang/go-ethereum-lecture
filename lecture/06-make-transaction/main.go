package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
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

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}
	// 1 ether = 1000000000000000000 wei
	amount := big.NewInt(100000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("./wallet/UTC--2023-01-06T07-27-55.500981000Z--c405d88eef718be968d466b3fadb089c8657bba4")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	//https://goerli.etherscan.io/tx/0xe2736d7c779026de25f0ce6dd813581fe68fcfb050a00eb44e7a42ae5b3c3a46
}
