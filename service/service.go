package service

import (
	"fmt"
	"os"
	"time"

	"github.com/BSaunders95/accounts-statistics-tool/config"
	"github.com/BSaunders95/accounts-statistics-tool/db"
	log "github.com/sirupsen/logrus"
)
const (
    layoutGB  = "2 January 2006"
)
type Service interface {
	GetNumberOfCICReports(dataDescription string)
}

type Impl struct {
	transactionClient db.TransactionClient
}

func NewService(cfg *config.Config) Service {

	return &Impl{
		transactionClient: db.NewTransactionDatabaseClient(cfg),
	}
}

func (s *Impl) GetNumberOfCICReports(dataDescription string) {

	transactions, err := s.transactionClient.GetAccountsTransactions(dataDescription)
	if err != nil {
		log.Error(fmt.Sprintf("Something went horribly wrong, oh no! %s", err))
		os.Exit(1)
	}

	closedTransactions := len(*transactions)
	acceptedTransactions := 0
	rejectedTransactions := 0
	januaryTransactions := 0
	februaryTransactions := 0
	marchTransactions := 0
	aprilTransactions := 0
	mayTransactions := 0
	juneTransactions := 0
	julyTransactions := 0
	augustTransactions := 0
	septemberTransactions := 0
	octoberTransactions := 0
	novemberTransactions := 0
	decemberTransactions := 0
	withinYearTransactions := 0

	yearAgo := time.Now().AddDate(-1, 0, 0)

	//time.local := time.utc

	for _, transaction := range *transactions {

		if transaction.Data.Filings != nil {
			filing := transaction.Data.Filings[transaction.ID+"-1"]
			if filing.Status == "accepted" {
				if transaction.Data.Closed_at.After(yearAgo) {
					withinYearTransactions++
				}
				switch transaction.Data.Closed_at.Month() {
				case 1:
					januaryTransactions++
				case 2:
					februaryTransactions++
				case 3:
					marchTransactions++
				case 4:
					aprilTransactions++
				case 5:
					mayTransactions++
				case 6:
					juneTransactions++
				case 7:
					julyTransactions++
				case 8:
					augustTransactions++
				case 9:
					septemberTransactions++
				case 10:
					octoberTransactions++
				case 11:
					novemberTransactions++
				case 12:
					decemberTransactions++
				}
				acceptedTransactions++
			} else if filing.Status == "rejected" {
				rejectedTransactions++
			}
		}
	}

	log.Info(fmt.Sprintf("Number of closed transactions: %d", closedTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions: %d", acceptedTransactions))
	log.Info(fmt.Sprintf("Number of rejected transactions: %d", rejectedTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in January  : %d", januaryTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in February : %d", februaryTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in March    : %d", marchTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in April    : %d", aprilTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in May      : %d", mayTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in June     : %d", juneTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in July     : %d", julyTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in August   : %d", augustTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in September: %d", septemberTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in October  : %d", octoberTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in November : %d", novemberTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions in December : %d", decemberTransactions))
	log.Info(fmt.Sprintf("Number of accepted transactions from %s: %d", yearAgo.Format(layoutGB), withinYearTransactions))
	fmt.Sprintf("Number of accepted transactions from %s: %d", yearAgo.Format(layoutGB), withinYearTransactions)
}
