package workflow

import (
	"github.com/mmuoDev/wallet/gen/wallet"
	"github.com/mmuoDev/wallet/internal/db"
	"github.com/pkg/errors"
)

//CreateWalletFunc provides a functionality to create a wallet
type CreateWalletFunc func(req wallet.CreateWalletRequest) error

//CreateWallet creates a wallet
func CreateWallet(createWallet db.CreateWalletFunc) CreateWalletFunc {
	return func(req wallet.CreateWalletRequest) error {
		data := make(map[string]interface{})
		data["account_id"] = req.AccountId
		data["previous_balance"] = req.PreviousBalance
		data["current_balance"] = req.CurrentBalance
		data["created_at"] = req.CreatedAt
		data["updated_at"] = req.UpdatedAt
		if err := createWallet(data); err != nil {
			return errors.Wrap(err, "workflow - unable to create wallet")
		}
		return nil
	}
}
