package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {

	// Kitaptan alindi
	// Keyfile ile ilgili seyler dert

	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	// keystoreFile := "wallets/UTC--2021-11-26T12-53-25.856008993Z--0e6a21feffd26ddf746b6356f794d14515a0db1b"
	// keyJson, err := ioutil.ReadFile(keystoreFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(keyJson))

	// password := "secret"
	// ks := keystore.KeyStore{}
	// account, err := ks.Import(keyJson, password, "")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ksJson, _ := json.MarshalIndent(ks, "", "   ")
	// accountJson, _ := json.MarshalIndent(account, "", "   ")

	// fmt.Println(string(ksJson))
	// fmt.Println(string(accountJson))

	// account, err := ks.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	// value := big.NewInt(1000000000000000000) // in wei (1 eth)
	// gasLimit := uint64(21000)                // in units
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("merhaba yalan dunya")
}
