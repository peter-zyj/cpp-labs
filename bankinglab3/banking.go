package banking

import "fmt"

var Transactions []Transaction

//var Transactions = []Transaction{{CHECKING, 1}, {CHECKING, 2},
//	{SAVINGS, 2}}

const (
	CHECKING = iota
	SAVINGS
)

type Transaction struct {
	accountType int
	amount      int
}

func DepositChecking(amount uint) int {
	Transactions = append(Transactions, Transaction{CHECKING, int(amount)})
	checking, _ := GetBalance()
	return checking
}

func WithdrawalChecking(amount int) (int, error) {
	checking, _ := GetBalance()
	if amount > checking {
		return checking, fmt.Errorf("Insufficient Funds")
	}
	Transactions = append(Transactions, Transaction{CHECKING, int(-amount)})
	return checking - amount, nil
}

func DepositSavings(amount int) int {
	Transactions = append(Transactions, Transaction{SAVINGS, int(amount)})
	_, savings := GetBalance()
	return savings
}

func WithdrawalSavings(amount int) (int, error) {
	_, savings := GetBalance()
	if amount > savings {
		return savings, fmt.Errorf("Insufficient Funds")
	}
	Transactions = append(Transactions, Transaction{SAVINGS, int(-amount)})
	return savings - amount, nil
}

func GetBalance() (int, int) {
	var checking int = 0
	var savings int = 0

	for _, item := range Transactions {
		if item.accountType == CHECKING {
			checking += item.amount
			continue
		}
		savings += item.amount
	}
	return checking, savings
}
