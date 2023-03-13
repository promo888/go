package v1

import (
	"errors"
	"fmt"
	"time"

	"github.com/vmihailenco/msgpack"
	mp "github.com/vmihailenco/msgpack/v5"
)

//"encoding/gob"

// type MessageType int

// const (
// 	MsgType MessageType = iota
// 	Message
// 	Transaction
// 	Contract
// 	Asset //new asset type, assetTx, assetFee, assetState
// 	Block
// 	Vote
// 	Account
// 	// AccountState
// 	// BlockState
// 	// ContractState
// 	// AsssetState
// )

// var MessageTypes = [...]MessageType{Id, Message, Transaction, Contract, Asset, Block, Vote, Account}

type Header struct { //updated by a miners pool by consensus
	Timestamp int64  `msgpack:"timestamp"` //unix timestamp utc - prepended by accepting miner
	Hash      string `msgpack:"hash"`      //msg hash  - hash of a msgpack encoded message
	//RootHash  string `msgpack:"roothash"`  //root block hash - genesis is default, fork or new token block - calculated trie
	BlockNum  int64  `msgpack:"blocknum"`  //block number - an indexer prepended after consensus and block approved
	Signature []byte `msgpack:"signature"` //signature of a proposer miner node

}

type Message struct {
	Header Header
	Data   []byte `msgpack:"data"` // tx, contract, asset, block, vote, account, etc
}

//todo: ? include last inputs/output in tx data
type Transfer struct {
	Timestamp int64   `msgpack:"timestamp"` //unix timestamp utc - nonce for tx
	From      string  `msgpack:"from"`      //from account, contract, or company
	To        string  `msgpack:"to"`        //to account, contract, or company
	Asset     string  `msgpack:"asset"`     //asset type: 1 - default blockchain asset, enumerated asset types (e.g. 2 - gold, 3 - silver, etc)
	Amount    float32 `msgpack:"amount"`    //amount of asset
	Fee       float32 `msgpack:"fee"`       //fee for tx
	//AssetFee string  `msgpack:"assetFee"` //fee asset type: hash(xcoin) - default blockchain asset, enumerated asset types (e.g. 2 - hash(gold), 3 - hash(silver), etc)
	Description string `msgpack:"description"` //description of tx - todo limit bytes by the config
	//RootHash  string `msgpack:"roothash"`  //root block hash - genesis is default, fork or new token block - calculated trie
	Signature []byte `msgpack:"signature"` //signature of a proposer miner node
	///////////////////// todo: prepend consensus % derived from account setup (e.g. 2 of 3, 3 of 5, or 55%, 65% etc for multisig accounts)
	// Signatures [][]byte `msgpack:"signatures"` //multiple signatures by a predefined account consensus
}

type Transaction struct {
	Message Message
}

// type Vote struct {
// 	Message Message
// }

//type Transaction Message
type Vote Message
type Contract Message
type Asset Message
type Account Message
type AccountState Message
type BlockState Message
type ContractState Message
type AsssetState Message

type Block struct {
	Header       Header
	PreviousHash string `msgpack:"prevhash"`
	Messages     []Message
	Votes        []Vote
	//Miner proposal and custom, board msgs to include in the block ?, or stroring in a helper db ?

}

func Echo() {
	fmt.Println("Echo v1 test")
}

// MarshalMsgPack marshals a Go object to MessagePack format
func MarshalMsgPack(obj interface{}) ([]byte, error) {
	msg, err := msgpack.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// UnmarshalMsgPack unmarshals a MessagePack message to a Go object
func UnmarshalMsgPack(msg []byte) (interface{}, error) {
	var obj interface{}
	err := mp.Unmarshal(msg, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// func (trf *Transfer) NewTransfer(from string, to string, asset string, amount float32, fee float32, description string) *Transfer {
// 	return &Transfer{
// 		Timestamp:   time.Now().Unix(),
// 		From:        from,
// 		To:          to,
// 		Asset:       asset,
// 		Amount:      amount,
// 		Fee:         fee,
// 		Description: description,
// 	}
// }

func NewTransaction(from string, to string, asset string, amount float32, fee float32, description string) (*Transaction, error) {
	// Validate input parameters - todo validation by a message type
	if from == "" {
		return nil, errors.New("from address cannot be empty")
	}
	if to == "" {
		return nil, errors.New("to address cannot be empty")
	}
	if asset == "" {
		return nil, errors.New("asset name cannot be empty")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	if fee <= 0 {
		return nil, errors.New("fee must be greater than zero")
	}

	transfer := &Transfer{
		Timestamp:   time.Now().Unix(),
		From:        from,
		To:          to,
		Asset:       asset,
		Amount:      amount,
		Fee:         fee,
		Description: description,
	}
	data, err := MarshalMsgPack(*transfer)
	if err != nil {
		return nil, err
	}
	// message := &Message{
	// 	Header: Header{},
	// 	Data:   data,
	// }

	message := &Transaction{
		Message: Message{
			Header: Header{},
			Data:   data,
		},
	}

	return message, nil
}

func (tx *Transaction) Serialize() ([]byte, error) {
	return MarshalMsgPack(tx)
}

func (tx *Transaction) Deserialize(data []byte) (*Transaction, error) {
	var transaction Transaction
	err := mp.Unmarshal(data, &transaction)
	if err != nil {
		return &transaction, err
	}
	return &transaction, nil
}

// func (t *Transaction) Serialize() ([]byte, error) {
// 	var result bytes.Buffer
// 	encoder := gob.NewEncoder(&result)
// 	err := encoder.Encode(t)
// 	return result.Bytes(), err
// }

// func DeserializeTransaction(data []byte) (Transaction, error) {
// 	var transaction Transaction
// 	decoder := gob.NewDecoder(bytes.NewReader(data))
// 	err := decoder.Decode(&transaction)
// 	return transaction, err
// }

// type MessageType, Message, AssetType, TransactionType, BlockType, VoteType int

// const (
// 	MessageType Id = iota
// 	Message
// 	Transaction
// 	Asset
// 	Block
// 	Vote
// 	//internal rpc messages ?
// 	Sync
// 	Contract
// 	Validation
// 	Delegation
// 	Execution

// )
// //todo: msgType : regular, contract, company, rpc, etc
// var MessageTypes = [...]MessageType{Id, Message, Transaction, Asset, Block, Vote, Sync, Contract, Validation, Delegation, Execution}

// const (
// 	TransactionType Id = iota
// 	Transfer
// 	Delegate
// 	Contract
// 	Lock
// )
// var TransactionTypes = [...]TransactionType{Id, Transfer, Delegate, Contract, Lock}

// const (
// 	AssetType Id = iota
// 	New //new asset, new entity/company, new token with a given economy/rule set
// 	Transfer //transfer asset to another account, transfer company to another account
// 	Delegate //delegate asset to another account for staking, voting, etc
// 	Contract //contract should accept a given asset type
// 	Lock //lock asset for a given period of time
// )
// var AssetTypes = [...]AssetType{Id, New, Transfer, Delegate, Contract, Lock}
