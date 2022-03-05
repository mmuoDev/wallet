package workflow

import (
	"strconv"

	"github.com/mmuoDev/core-proto/gen/wallet"
	"github.com/mmuoDev/wallet/internal/db"
	"github.com/pkg/errors"
)

//CreateWalletFunc provides a functionality to create a wallet
type CreateWalletFunc func(req wallet.CreateWalletRequest) (wallet.CreateWalletResponse, error)

//RetrieveWalletFunc provides a functionality to retrieve wallet by account id
type RetrieveWalletFunc func(req wallet.RetrieveWalletRequest) (wallet.RetrieveWalletResponse, error)

//UpdateWalletFunc provides a functionality to update a wallet
type UpdateWalletFunc func(req wallet.UpdateWalletRequest) error

//CreateWallet creates a wallet
func CreateWallet(createWallet db.CreateWalletFunc) CreateWalletFunc {
	return func(req wallet.CreateWalletRequest) (wallet.CreateWalletResponse, error) {
		lastId, err := createWallet(req)
		if err != nil {
			return wallet.CreateWalletResponse{}, errors.Wrap(err, "workflow - unable to create wallet")
		}
		return wallet.CreateWalletResponse{Id: lastId}, nil
	}
}

//UpdateWallet updates a wallet
func UpdateWallet(updateWallet db.UpdateWalletFunc) UpdateWalletFunc {
	return func(req wallet.UpdateWalletRequest) error {
		if err := updateWallet(req); err != nil {
			return errors.Wrap(err, "unable to update wallet")
		}
		return nil
	}
}

//RetrieveWallet retrieves a wallet
func RetrieveWallet(retrieveWallet db.RetrieveWalletByAccountIdFunc) RetrieveWalletFunc {
	return func(req wallet.RetrieveWalletRequest) (wallet.RetrieveWalletResponse, error) {
		accID := strconv.FormatInt(int64(req.AccountId), 10)
		w, err := retrieveWallet(accID)
		if err != nil {
			return wallet.RetrieveWalletResponse{}, errors.Wrap(err, "workflow - unable to retrieve wallet")
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
