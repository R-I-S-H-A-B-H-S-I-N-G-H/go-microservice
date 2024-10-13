package services

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/aws_util"
	"time"

	"github.com/google/uuid"
)

type WalletService struct {
	// Potentially other fields related to wallet services
}

type Wallet struct {
	Id               string         `json:"id"`
	Transactions     []*Transaction `json:"transactions"`
	Tags             []*Tag         `json:"tags"`
	Date             string         `json:"date"`
	DefaultIncomeTag *Tag           `json:"defaultIncomeTag"`
}

// Transaction structure
type Transaction struct {
	Id      string  `json:"id"`
	Amount  float64 `json:"amount"`
	Date    string  `json:"date"`
	Type    string  `json:"type"`
	Tag     *Tag    `json:"tag"`
	Comment string  `json:"comment"`
}

// Tag structure
type Tag struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	Type   string  `json:"type"`
}

// create an enum for transaction type
const (
	CREDIT = "CREDIT"
	DEBIT  = "DEBIT"
)

//how to use it

func CreateNewWallet() *Wallet {
	return &Wallet{
		Id:               uuid.NewString(),
		Transactions:     []*Transaction{},
		Tags:             []*Tag{},
		Date:             time.Now().Format(time.RFC3339),
		DefaultIncomeTag: nil,
	}
}

// Adding transactions
func (w *Wallet) addTransaction(transaction *Transaction) {
	w.Transactions = append(w.Transactions, transaction)
}

// Adding Tag
func (w *Wallet) addTag(tag *Tag) {
	w.Tags = append(w.Tags, tag)
}

func createNewTransaction(amount float64, comment string, date string, type_ string, tag *Tag) *Transaction {
	return &Transaction{
		Id:      uuid.NewString(),
		Amount:  amount,
		Date:    date,
		Type:    type_,
		Tag:     tag,
		Comment: comment,
	}
}

func createNewTag(name string, amount float64, date string, type_ string) *Tag {
	return &Tag{
		Id:     uuid.NewString(),
		Name:   name,
		Amount: amount,
		Date:   date,
		Type:   type_,
	}
}

func GetWalletList() []*Wallet {
	return []*Wallet{
		GetNewWallet(),
	}
}

func GetNewWallet() *Wallet {
	wallet := CreateNewWallet()
	tag := createNewTag("Income", 99999999999999, time.Now().Format(time.RFC3339), CREDIT)
	wallet.addTag(tag)

	return wallet
}

func SyncWalletToS3(userId string, walletPayload string) (string, error) {
	return aws_util.UploadStrDataToS3(userId+"/wallet.json", walletPayload)
}
