package db

import (
	"fmt"
	"strconv"

	"github.com/mmuoDev/wallet/gen/wallet"
	pg "github.com/mmuoDev/wallet/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	walletTable = "wallet"
)

//Wallet represents a row of wallet data
type Wallet struct {
	AccountID       int `json:"account_id"`
	PreviousBalance int `json:"previous_balance"`
	CurrentBalance  int `json:"current_balance"`
}

//CreateWalletFunc provides functionality to create wallet
type CreateWalletFunc func(data map[string]interface{}) error

//RetrieveWalletByAccountIdFunc provides functionality to retrieve a wallet by account id
type RetrieveWalletByAccountIdFunc func(accID string) (Wallet, error)

//UpdateWalletFunc returns a functionality to update a wallet
type UpdateWalletFunc func(req wallet.UpdateWalletRequest) error

//CreateWallet creates a wallet on db
func CreateWallet(dbConnector pg.Connector) CreateWalletFunc {
	return func(data map[string]interface{}) error {
		_, err := dbConnector.Insert(walletTable, data)
		if err != nil {
			return errors.Wrap(err, "db - unable to insert record")
		}
		return nil
	}
}

//UpdateWallet updates a wallet
func UpdateWallet(dbConnector pg.Connector) UpdateWalletFunc {
	return func(req wallet.UpdateWalletRequest) error {
		query := fmt.Sprintf("UPDATE %s SET previous_balance = ?, current_balance = ? WHERE account_id = ?", walletTable)
		var params []interface{}
		pp := append(params, req.PreviousBalance, req.CurrentBalance, req.AccountId)
		if err := dbConnector.Update(query, pp); err != nil {
			return errors.Wrap(err, "unable to update wallet")
		}
		return nil
	}
}

//RetrieveWalletByAccountId retrieves a wallet by account id
func RetrieveWalletByAccountId(dbConnector pg.Connector) RetrieveWalletByAccountIdFunc {
	return func(accID string) (Wallet, error) {
		accId := "account_id"
		prevBalance := "previous_balance"
		curBalance := "current_balance"
		query := fmt.Sprintf("SELECT %s, %s, %s FROM %s where account_id = ?", accId, prevBalance, curBalance, walletTable)
		var params []interface{}
		pp := append(params, accID)
		_, err := dbConnector.Select(query, pp, &accID, &prevBalance, &curBalance)
		if err != nil {
			return Wallet{}, errors.Wrap(err, "no row found")
		}
		pb, err := stringToInt(prevBalance)
		if err != nil {
			return Wallet{}, errors.Wrap(err, "unable to convert previous balance to int")
		}
		cb, err := stringToInt(curBalance)
		if err != nil {
			return Wallet{}, errors.Wrap(err, "unable to convert current balance to int")
		}
		aId, err := stringToInt(accId)
		if err != nil {
			return Wallet{}, errors.Wrap(err, "unable to convert current balance to int")
		}
		w := Wallet{
			AccountID:       aId,
			PreviousBalance: pb,
			CurrentBalance:  cb,
		}
		return w, nil
	}
}

//stringToInt converts a string to int
func stringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
