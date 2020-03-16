package reverse

import "testing"

func makeRange(max int) []int {
	a := make([]int, max+1)
	for i := range a {
		a[i] = i
	}
	return a
}

func BenchmarkRevese1(b *testing.B) {
	arr := makeRange(200000000)
	b.ResetTimer()
	ReverseArray(arr)
}

func BenchmarkRevese2(b *testing.B) {
	arr := makeRange(400000000)
	b.ResetTimer()
	ReverseArray(arr)
}
