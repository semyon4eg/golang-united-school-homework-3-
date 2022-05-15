package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	
	keys := make([]int, 0, len(input))

	result = make([]string, 0, len(input))

	for key := range input {
		keys = append (keys, key)
	}

	sort.Ints(keys)

	for k := range keys {
		result = append(result, input[k])
	}

	return result
}
