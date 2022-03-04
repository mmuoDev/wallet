package workflow

import (
	"strconv"

	"github.com/mmuoDev/wallet/gen/wallet"
	"github.com/mmuoDev/wallet/internal/db"
	"github.com/pkg/errors"
)

//CreateWalletFunc provides a functionality to create a wallet
type CreateWalletFunc func(req wallet.CreateWalletRequest) error

//RetrieveWalletFunc provides a functionality to retrieve wallet by account id
type RetrieveWalletFunc func(req wallet.RetrieveWalletRequest) (wallet.RetrieveWalletResponse, error)

//UpdateWalletFunc provides a functionality to update a wallet
type UpdateWalletFunc func(req wallet.UpdateWalletRequest) error

//CreateWallet creates a wallet
func CreateWallet(createWallet db.CreateWalletFunc) CreateWalletFunc {
	return func(req wallet.CreateWalletRequest) error {
		data := make(map[string]interface{})
		data["account_id"] = req.AccountId
		data["previous_balance"] = req.PreviousBalance
		data["current_balance"] = req.CurrentBalance
		if err := createWallet(data); err != nil {
			return errors.Wrap(err, "workflow - unable to create wallet")
		}
		return nil
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
		aID, err := stringToInt(w.AccountID)
		if err != nil {
			return wallet.RetrieveWalletResponse{}, errors.Wrap(err, "workflow - unable to convert accountId to int")
		}
		res := wallet.RetrieveWalletResponse{
			AccountId:       int32(aID),
			PreviousBalance: int32(w.PreviousBalance),
			CurrentBalance:  int32(w.CurrentBalance),
		}
		return res, nil
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
