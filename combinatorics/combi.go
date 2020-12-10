package combinatorics

// CombinationsRecursive0 is a naive recursive implementation.
func CombinationsRecursive0(begin, end, k int) [][]int {
	if k == 0 {
		pattern := []int{}
		return [][]int{pattern}
	}

	ans := [][]int{}
	for num := begin; num < end-k+1; num++ {
		childPatterns := CombinationsRecursive0(num+1, end, k-1)
		for _, childPattern := range childPatterns {
			pattern := append([]int{num}, childPattern...)
			ans = append(ans, pattern)
		}
	}

	return ans
}

// CombinationsRecursive1 uses the same memory space for each combination.
func CombinationsRecursive1(n, k int, f func([]int)) {
	pattern := make([]int, k)

	var body func(pos, begin int)
	body = func(pos, begin int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := begin; num < n+pos-k+1; num++ {
			pattern[pos] = num
			body(pos+1, num+1)
		}
	}
	body(0, 0)
}

// CombinationsRecursive2 does not pass the beginning of the available number
// range to recursive call. It refers the generating combination instead.
func CombinationsRecursive2(n, k int, f func([]int)) {
	pattern := make([]int, k+1)
	pattern[0] = -1

	var body func(pos int)
	body = func(pos int) {
		if pos == k+1 {
			f(pattern[1:])
			return
		}

		for num := pattern[pos-1] + 1; num < n+pos-k; num++ {
			pattern[pos] = num
			body(pos + 1)
		}
	}
	body(1)
}

// CombinationsWithCarrying0 decides the next digit of combination
// to increment not by recursive calls but by the previous combination
// directly.
func CombinationsWithCarrying0(n, k int, f func([]int)) {
	pattern := make([]int, k)
	for i := range pattern {
		pattern[i] = i
	}

	for {
		f(pattern)

		pos := k - 1
		for pos >= 0 {
			oldNum := pattern[pos]

			if oldNum == n+pos-k {
				// carry
				pos--
				continue
			}

			// increment
			pattern[pos]++
			break
		}
		if pos == -1 {
			break
		}

		// replace the numbers of carried digits
		for pos++; pos < k; pos++ {
			pattern[pos] = pattern[pos-1] + 1
		}
	}
}

// CombinationsWithCarrying1 integrates the loop for increment and carrying,
// which goes from right digit to left digit, and the one for setting rest
// values, which goes from left digit to right digit.
func CombinationsWithCarrying1(n, k int, f func([]int)) {
	pattern := make([]int, k)
	for i := range pattern {
		pattern[i] = i
	}

	pos := k
	for pos > -1 {
		if pos == k {
			f(pattern)
			pos--
			continue
		}

		// carry
		oldNum := pattern[pos]
		if oldNum == n+pos-k {
			pattern[pos] = -1
			pos--
			continue
		}

		if oldNum == -1 {
			// replace the number of the carried digit
			pattern[pos] = pattern[pos-1] + 1
		} else {
			// increment
			pattern[pos]++
		}
		pos++
	}
}
