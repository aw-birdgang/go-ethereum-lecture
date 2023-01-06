package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(pData))

	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(puData))

	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
