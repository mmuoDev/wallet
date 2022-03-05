package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"

	pb "github.com/mmuoDev/core-proto/gen/wallet"
	"github.com/mmuoDev/wallet/internal/server"
	pg "github.com/mmuoDev/wallet/pkg/postgres"
	"github.com/stretchr/testify/assert"
)

//postgresProvider mocks postgres
func postgresProvider() *pg.Connector {
	return &pg.Connector{}
}

func TestCreateWalletServiceWorkAsExpected(t *testing.T) {
	isDBInvoked := false
	mockDB := func(o *server.Option) {
		o.CreateWallet = func(req pb.CreateWalletRequest) (int64, error) {
			isDBInvoked = true
			t.Run("Data is as expected", func(t *testing.T) {
				assert.Equal(t, int32(926592), req.AccountId)
				assert.Equal(t, int32(100), req.PreviousBalance)
				assert.Equal(t, int32(100), req.CurrentBalance)
			})
			return 1, nil
		}
	}
	opts := []server.OptionalArg{
		mockDB,
	}
	ap := server.New(&pg.Connector{}, opts...)
	req := &pb.CreateWalletRequest{AccountId: 926592, PreviousBalance: 100, CurrentBalance: 100}
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
	var dbRes pb.RetrieveWalletResponse
	fileToStruct(filepath.Join("testdata", "retrieve_wallet_db_response.json"), &dbRes)
	mockDB := func(o *server.Option) {
		o.RetrieveWallet = func(accID string) (pb.RetrieveWalletResponse, error) {
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
		assert.Equal(t, res.GetAccountId(), int32(926592))
	})
	t.Run("Previous balance is as expected", func(t *testing.T) {
		assert.Equal(t, res.GetPreviousBalance(), int32(1000))
	})
	t.Run("Current balance is as expected", func(t *testing.T) {
		assert.Equal(t, res.GetCurrentBalance(), int32(500))
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
