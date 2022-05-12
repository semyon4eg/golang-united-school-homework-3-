package homework

func sortMapValues(input map[int]string) (result []string) {
	
	var min int

	for i := 0; i < len(input); i++ {

		for key := range input {
			min = key
			if key < min {
				min = key
			}
		}

		result = append(result, input[min])
		delete(input, min)
	}

	return result
}
