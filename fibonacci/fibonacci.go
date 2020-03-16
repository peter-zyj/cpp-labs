package fibonacci

import "fmt"

func ValidateFibonacci(sequence []int) (bool, error) {
	if len(sequence) < 3 {
		return false, fmt.Errorf("sequence to short")
	}

	if (sequence[0] != 0) || (sequence[1] != 1) {
		return false, fmt.Errorf("invalid start")
	}

	for i := 2; i < len(sequence); i++ {
		expected := sequence[i-1] + sequence[i-2]
		if sequence[i] != expected {
			return false, fmt.Errorf("invalid sequence")
		}
	}
	return true, nil
}
