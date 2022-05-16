package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here

	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	
	sort.Ints(keys)

	values := make([]string,0,len(input))
	for _, k := range keys {
		values = append(values, input[k])
	}

	return values
}
