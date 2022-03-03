package db

import (
	pg "github.com/mmuoDev/wallet/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	walletTable = "wallet"
)

//CreateWalletFunc provides functionality to create wallet
type CreateWalletFunc func(data map[string]interface{}) error

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
