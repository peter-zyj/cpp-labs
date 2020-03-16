package banking

import (
	"testing"
)

func TestDepositCheckingIfDeposit10(t *testing.T) {
	Transactions = Transactions[0:]
	expected := 10
	const START = 10
	result := DepositChecking(START)
	if result != expected {
		t.Error("invalid checking balance")
	}
}

func TestWithdrawalCheckingIfWithdrawal10(t *testing.T) {
	Transactions = Transactions[0:]
	expected := 0
	const START = 10
	const WITHDRAWAL = 10
	result := DepositChecking(START)
	result, err := WithdrawalChecking(WITHDRAWAL)
	if (result != expected) || (err != nil) {
		t.Error("invalid checking balance")
	}
}

func TestWithdrawalCheckingIfWithdrawalError(t *testing.T) {
	Transactions = Transactions[0:]
	const EXPECTED = 10
	const START = 10
	const WITHDRAWAL = 15
	DepositChecking(START)
	result, err := WithdrawalChecking(WITHDRAWAL)
	if (EXPECTED != result) || (err == nil) {
		t.Error("invalid  balance")
	}
}

func TestDepositSavingsIfDeposit10(t *testing.T) {
	Transactions = Transactions[0:]
	expected := 10
	const START = 10
	result := DepositSavings(START)
	if result != expected {
		t.Error("invalid savings balance")
	}
}

func TestWithdrawalSavingsIfWithdrawal10(t *testing.T) {
	Transactions = Transactions[0:]
	expected := 0
	const START = 10
	const WITHDRAWAL = 10
	result := DepositSavings(START)
	result, err := WithdrawalSavings(WITHDRAWAL)
	if (result != expected) || (err != nil) {
		t.Error("invalid savings balance")
	}
}

func TestWithdrawalSavingsIfWithdrawalError(t *testing.T) {
	Transactions = Transactions[0:]
	const EXPECTED = 10
	const START = 10
	const WITHDRAWAL = 15
	DepositSavings(START)
	result, err := WithdrawalSavings(WITHDRAWAL)
	if (EXPECTED != result) || (err == nil) {
		t.Error("invalid savings")
	}
}

func TestGetBalanceIfBalances10(t *testing.T) {
	Transactions = Transactions[0:]
	const EXPECTED = 10
	const START = 10
	DepositChecking(START)
	DepositSavings(START)
	checking, savings := GetBalance()
	if (checking != EXPECTED) || (savings != EXPECTED) {
		t.Error("invalid balances")
	}
}
