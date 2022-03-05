package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"

	pb "github.com/mmuoDev/wallet/gen/wallet"
	"github.com/mmuoDev/wallet/internal/server"
	"github.com/mmuoDev/wallet/internal/db"
	pg "github.com/mmuoDev/wallet/pkg/postgres"
	"github.com/stretchr/testify/assert"
)

//postgresProvider mocks postgres
func postgresProvider() *pg.Connector {
	return &pg.Connector{}
}

func TestCreateWalletServiceWorkAsExpected(t *testing.T) {
	expectedAccount := int32(926592)
	isDBInvoked := false
	mockDB := func(o *server.Option) {
		o.CreateWallet = func(data map[string]interface{}) error {
			isDBInvoked = true
			t.Run("Data is as expected", func(t *testing.T) {
				assert.Equal(t, expectedAccount, data["account_id"])
			})
			return nil
		}
	}
	opts := []server.OptionalArg{
		mockDB,
	}
	ap := server.New(&pg.Connector{}, opts...)
	req := &pb.CreateWalletRequest{AccountId: 926592, PreviousBalance: 100}
	_, err := ap.CreateWallet(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("DB is invoked", func(t *testing.T) {
		assert.True(t, isDBInvoked)
	})
}

func TestUpdateWalletServiceWorkAsExpected(t *testing.T) {
	isDBInvoked := false
	mockDB := func(o *server.Option) {
		o.UpdateWallet = func(req pb.UpdateWalletRequest) error {
			isDBInvoked = true
			t.Run("Data is as expected", func(t *testing.T) {
				assert.Equal(t, req.CurrentBalance, int32(1500))
			})
			return nil 
		}
	}
	opts := []server.OptionalArg{
		mockDB,
	}
	ap := server.New(&pg.Connector{}, opts...)
	req := &pb.UpdateWalletRequest{AccountId: 926592, PreviousBalance: 100, CurrentBalance: 1500}
	_, err := ap.UpdateWallet(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("DB is invoked", func(t *testing.T) {
		assert.True(t, isDBInvoked)
	})
}
func TestRetrieveWalletServiceWorkAsExpected(t *testing.T) {
	isDBInvoked := false
	expectedAccount := int32(926592)
	var dbRes db.Wallet
	fileToStruct(filepath.Join("testdata", "retrieve_wallet_db_response.json"), &dbRes)
	mockDB := func(o *server.Option) {
		o.RetrieveWallet = func(accID string) (db.Wallet, error) {
			isDBInvoked = true
			return dbRes, nil 
		}
	}
	opts := []server.OptionalArg{
		mockDB,
	}
	ap := server.New(&pg.Connector{}, opts...)
	req := &pb.RetrieveWalletRequest{AccountId: 926592}
	res, err := ap.RetrieveWallet(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Account is as expected", func(t *testing.T) {
		assert.Equal(t, res.GetAccountId(), expectedAccount)
	})
	t.Run("DB is invoked", func(t *testing.T) {
		assert.True(t, isDBInvoked)
	})
}

// fileToStruct reads a json file to a struct
func fileToStruct(filepath string, s interface{}) io.Reader {
	bb, _ := ioutil.ReadFile(filepath)
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb)
}
