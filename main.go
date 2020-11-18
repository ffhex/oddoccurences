package main

func Solution(A []int) int {

	if len(A) % 2 == 0 {
		return 0
	}

	m := make(map[int]int)
	for _, i := range A {
		if _, exists := m[i]; !exists {
			m[i] = 1
		} else {
			m[i]++
		}
	}

	for k, v := range m {
		if v % 2 != 0 {
			return k
		}
	}

	return 0
}
