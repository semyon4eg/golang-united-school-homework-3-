package homework

func sortMapValues(input map[int]string) (result []string) {
	
	var min int

	for i := 0; i < len(input); i++ {

		for k := range input {
			min = k
			break
		}

		for key := range input {
			if key < min {
				min = key
			}
		}

		result = append(result, input[min])
		delete(input, min)
	}

	return result
}
