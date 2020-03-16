package badunittests

import "testing"

var dataitems = []int{1, 2, 3}

func TestFuncA(t *testing.T) {

	TestFuncB(t)

	total := 0
	for _, item := range dataitems {
		total += item
		if item < 0 {
			t.Error()
		}
	}

	if total != 6 {
		t.Error()
	}

	dataitems = append(dataitems, 4)
}

func TestFuncB(t *testing.T) {
	result := display(12, 5)
	if result == false {
		t.Error()
	}
}

func TestDivideThrice(t *testing.T) {
	result, _ := DivideThrice(81, 3)
	if result != 3 {
		t.Error()
	}
}
