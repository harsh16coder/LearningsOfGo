package pointersbasedtests

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit money", func(t *testing.T) {
		wallets := Wallet{}
		wallets.Deposit(Bitcoin(20))
		got := wallets.Balance()
		want := Bitcoin(20)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw money", func(t *testing.T) {
		wallets := Wallet{balance: Bitcoin(30)}
		wallets.Withdraw(20)
		got := wallets.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Insufficient money", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		asserError(t, err, ErrBalanceStatement)
		assertBalance(t, wallet.balance, Bitcoin(20))
	})
}

func asserError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("Expected error but didn't got")
	}
	if got == want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func assertBalance(t testing.TB, got Bitcoin, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
