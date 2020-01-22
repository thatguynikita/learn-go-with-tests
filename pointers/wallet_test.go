package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(15))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	
	if got != want {
		t.Errorf("got %s but expected %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q but expected %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("got an error but didn't want one - %s", got.Error())
	}
}