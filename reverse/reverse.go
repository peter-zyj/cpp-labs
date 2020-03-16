package reverse

func ReverseArray(array []int) []int {
	temp := 0
	start := 0
	end := len(array) - 1
	for start < end {
		temp = array[start]
		array[start] = array[end]
		array[end] = temp
		start++
		end--
	}
	return array
}
