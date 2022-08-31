package wallet

import (
	"testing"
)

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("Wallet  Deposites", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposite(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Wallet withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}

		wallet.Withdraw(25)

		assertBalance(t, wallet, Bitcoin(25))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrinsufficientFunds)

	})

}
