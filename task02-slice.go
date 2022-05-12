package homework

func reverse(input []int64) (result []int64) {
	
	var reversed []int64

	length := len(input)

	for i := 0; i < length; i++ {
		reversed = append(reversed, input[length - 1 - i])
	}

	return reversed
}
