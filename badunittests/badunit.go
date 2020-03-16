package badunittests

import "fmt"

func Cubed(data int) int {
	result := data * data * data
	display(result, data)
	return result
}

func display(result int, value int) bool {
	fmt.Printf("%d=%d+%d+%d\n", result,
		value, value, value)
	return true
}

func DivideThrice(base float32,
	divide float32) (float32, error) {
	if divide == 0 {
		return 0, fmt.Errorf("Divide by zero")
	}
	return (base / divide / divide / divide), nil
}
