package main

// import (
// 	"fmt"

// 	"github.com/tyler-smith/go-bip32"
// 	"github.com/tyler-smith/go-bip39"
// )

// func main() {
// 	// generate a new mnemonic with 24 words
// 	entropy, err := bip39.NewEntropy(256)
// 	if err != nil {
// 	}
// 	mnemonic, err := bip39.NewMnemonic(entropy)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Mnemonic:", mnemonic)

// 	// generate a new seed from the mnemonic
// 	seed := bip39.NewSeed(mnemonic, "")

// 	// derive a master key from the seed
// 	masterKey, err := bip32.NewMasterKey(seed)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// derive an Algorand account from the master key
// 	path := "m/44'/283'/0'/0/0" // the standard derivation path for Algorand accounts
// 	algorandKey, err := masterKey.NewChildKey(bip32.MustParsePath(path))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// generate an Algorand address from the public key
// 	address := algorandKey.PublicKey().Address()
// 	fmt.Println("Address:", address)
// }

// 

import (
	"fmt"
	"log"

	"github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	mnemonic := "engage clever judge nice couple benefit glare material enact length perfect scene multiply siren provide immune donate pair fence velvet pencil polar funny maple"
	//mnemonic := "clever judge engage nice couple benefit glare material enact length perfect scene multiply siren provide immune donate pair fence velvet pencil polar funny maple"

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/777'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947 "m/44'/60'/0'/0/0"

	path = hdwallet.MustParseDerivationPath("m/44'/777'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
} 