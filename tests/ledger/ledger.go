package main

import (
	"errors"
	"fmt"
)

type Asset struct {
	Symbol  string
	Balance int
}

type Account struct {
    Name   string
    Assets map[string]*Asset
}

type Transaction struct {
	From   *Account
	To     *Account
	Asset  *Asset
	Amount int
}

type Ledger struct {
	Transactions []*Transaction
}

func (l *Ledger) AddTransaction(tx *Transaction) {
	l.Transactions = append(l.Transactions, tx)
}

func (a *Account) AddAsset(asset *Asset) {
	if a.Assets == nil {
		a.Assets = make(map[string]*Asset)
	}
	a.Assets[asset.Symbol] = asset
}

func (a *Account) GetAsset(symbol string) *Asset {
	return a.Assets[symbol]
}

func (a *Account) AddToAsset(symbol string, amount int) {
	asset := a.GetAsset(symbol)
	if asset != nil {
		asset.Balance += amount
	}
}

func (a *Account) SubFromAsset(symbol string, amount int) error {
	asset := a.GetAsset(symbol)
	if asset == nil {
		return errors.New("asset not found")
	}
	if asset.Balance < amount {
		return errors.New("not enough balance")
	}
	asset.Balance -= amount
	return nil
}

func (t *Transaction) Execute() error {
	err := t.From.SubFromAsset(t.Asset.Symbol, t.Amount)
	if err != nil {
		return err
	}
	t.To.AddToAsset(t.Asset.Symbol, t.Amount)
	return nil
}

func main() {
	alice := &Account{Name: "Alice"}
	bob := &Account{Name: "Bob"}
	ledger := &Ledger{}

	// Add assets to Alice's account
	alice.AddAsset(&Asset{Symbol: "BTC", Balance: 10})
	alice.AddAsset(&Asset{Symbol: "ETH", Balance: 100})

	// Create a transaction to transfer 1 BTC from Alice to Bob
	tx := &Transaction{
		From:   alice,
		To:     bob,
		Asset:  alice.GetAsset("BTC"),
		Amount: 1,
	}

	// Execute the transaction and add it to the ledger
	err := tx.Execute()
	if err != nil {
		fmt.Println("Transaction failed:", err)
		return
	}
	ledger.AddTransaction(tx)

	// Print Alice's and Bob's balances after the transaction
	for _, asset := range alice.Assets {
		fmt.Printf("Alice's %s balance: %d\n", asset.Symbol, asset.Balance)
	}
	for _, asset := range bob.Assets {
		fmt.Printf("Bob's %s balance: %d\n", asset.Symbol, asset.Balance)
	}

	// Print the ledger transactions
	fmt.Println("Ledger Transactions:")
	for _, tx := range ledger.Transactions {
		fmt.Printf("%s transferred %d %s to %s\n", tx.From.Name, tx.Amount, tx.Asset.Symbol, tx.To.Name)
		//fmt.Printf("%p transferred %d %s to %p\n", tx.From, tx.Amount, tx.Asset.Symbol, tx.To)

	}
}
