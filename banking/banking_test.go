package banking

import (
	"fmt"
	"testing"
)

func TestBanking(t *testing.T) {
	DepositChecking(10)
	DepositSavings(20)
	checking, savings := GetBalance()
	fmt.Println(checking, savings)
}
