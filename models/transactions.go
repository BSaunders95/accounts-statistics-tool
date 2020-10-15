package models

// Transaction describes a transaction database entity
type Transaction struct {
	ID        string `bson:"_id"`
	Data      TransactionData `bson:"data"`
}

type TransactionData struct {
	Description string                       `bson:"description"`
	Filings     map[string]TransactionFiling `bson:"filings"`
	Links       map[string]string            `bson:"links"`
	Status      string                       `bson:"status"`
}

type TransactionFiling struct {
	Type   string `bson:"type"`
	Status string `bson:"status"`
}
