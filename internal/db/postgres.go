package db

import (
	"strconv"

	"database/sql"

	"github.com/mmuoDev/core-proto/gen/wallet"
	pg "github.com/mmuoDev/wallet/pkg/postgres"
	"github.com/pkg/errors"
)

//CreateWallet creates a wallet on db
func CreateWallet(dbConnector pg.Connector) CreateWalletFunc {
	return func(req wallet.CreateWalletRequest) (int64, error) {
		query := `INSERT INTO wallet (
			account_id, previous_balance, current_balance
		) VALUES (
			$1, $2, $3
		) RETURNING id`
		res, err := dbConnector.DB.Exec(query, req.AccountId, req.PreviousBalance, req.CurrentBalance)
		if err != nil {
			return 0, errors.Wrap(err, "db - unable to insert record")
		}
		id, err := res.LastInsertId()
		if err != nil {
			return 0, errors.Wrap(err, "db - unable to retrieve last inserted id")
		}
		return id, nil

	}
}

//UpdateWallet updates a wallet
func UpdateWallet(dbConnector pg.Connector) UpdateWalletFunc {
	return func(req wallet.UpdateWalletRequest) error {
		query := `UPDATE wallet SET previous_balance = $1, current_balance = $2 WHERE account_id = $3`
		_, err := dbConnector.DB.Exec(query, req.PreviousBalance, req.CurrentBalance, req.AccountId)
		return err
	}
}

//RetrieveWalletByAccountId retrieves a wallet by account id
func RetrieveWalletByAccountId(dbConnector pg.Connector) RetrieveWalletByAccountIdFunc {
	return func(accID string) (wallet.RetrieveWalletResponse, error) {
		var w wallet.RetrieveWalletResponse
		query := `SELECT account_id, previous_balance, current_balance FROM wallet where account_id = $1`
		row := dbConnector.DB.QueryRow(query, accID)
		switch err := row.Scan(&w.AccountId, &w.PreviousBalance, &w.CurrentBalance); err {
		case sql.ErrNoRows:
			return w, errors.New("not found")
		case nil:
			return w, nil
		default:
			return w, errors.Wrap(err, "error retrieving data")
		}
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
