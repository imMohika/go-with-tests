package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		assert(t, got, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Errorf("got an error %q", err)
			return
		}
		got := wallet.Balance()
		want := Bitcoin(10)

		assert(t, got, want)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		err := wallet.Withdraw(Bitcoin(100))
		got := wallet.Balance()
		want := Bitcoin(20)

		assertError(t, err, ErrInsufficientFunds)
		assert(t, got, want)
	})
}
func assert(t testing.TB, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("wanted an error but didn't got any")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
