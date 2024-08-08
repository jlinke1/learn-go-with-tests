package pointers_test

import (
	"errors"
	"testing"

	"github.com/jlinke1/learn-go-with-tests/pointers"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		t.Parallel()
		wallet := pointers.Wallet{}

		wallet.Deposit(10)

		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		t.Parallel()
		wallet := pointers.Wallet{}
		wallet.Deposit(10)

		err := wallet.Withdraw(5)

		assertNoError(t, err)
		assertBalance(t, wallet, pointers.Bitcoin(5))
	})

	t.Run("withdraw too much", func(t *testing.T) {
		t.Parallel()
		wallet := pointers.Wallet{}

		err := wallet.Withdraw(10)
		assertBalance(t, wallet, pointers.Bitcoin(0))

		assertError(t, err, pointers.ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet pointers.Wallet, want pointers.Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("didn't expect an error, but got: %v", got)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()

	if err == nil {
		t.Fatalf("wanted an error but didn''t get one")
	}

	if !errors.Is(err, want) {
		t.Errorf("got %q, want %q", err, want)
	}
}
