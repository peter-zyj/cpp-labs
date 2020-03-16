package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncNil(t *testing.T) {
	_, err := ValidateFibonacci([]int{0, 1})
	assert.NotNil(t, err, "value is nil")
}

func TestFuncValidate(t *testing.T) {
	result, _ := ValidateFibonacci([]int{0, 1, 1, 5, 2, 3, 5})
	assert.True(t, result, "sequence not valid")
}
