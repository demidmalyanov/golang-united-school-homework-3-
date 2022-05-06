package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here

	m := map[int]string{2: "test1", 1: "test2", 4: "test3"}

	keys := make([]int, 0, len(input))
	for k := range m {
		keys = append(keys, k)
	}
	
	sort.Ints(keys)

	values := make([]string,0,len(input))
	for _, k := range keys {
		values = append(values, input[k])
	}

	return values
}
