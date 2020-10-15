package service

import (
	"fmt"
	"github.com/BSaunders95/accounts-statistics-tool/config"
	"github.com/BSaunders95/accounts-statistics-tool/db"
	log "github.com/sirupsen/logrus"
	"os"
)

type Service interface {
	GetNumberOfCICReports()
}

type Impl struct {
	transactionClient db.TransactionClient
}

func NewService(cfg *config.Config) Service {

	return &Impl{
		transactionClient: db.NewTransactionDatabaseClient(cfg),
	}
}

func (s *Impl) GetNumberOfCICReports() {

	transactions, err := s.transactionClient.GetAccountsTransactions()
	if err != nil {
		log.Error(fmt.Sprintf("Something went horribly wrong, oh no! %s", err))
		os.Exit(1)
	}

	closedTransactions := len(*transactions)
	acceptedTransactions := 0

	for _, transaction := range *transactions {

		if transaction.Data.Filings != nil {
			filing := transaction.Data.Filings[transaction.ID + "-1"]
			if filing.Status == "accepted" {
				acceptedTransactions++
			}
		}
	}

	log.Info(fmt.Sprintf("Number of closed transactions: %d", closedTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions: %d", acceptedTransactions))
}
