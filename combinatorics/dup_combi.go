package combinatorics

// DupCombinationsRecursive0 is a naive recursive implementation.
func DupCombinationsRecursive0(begin, end, k int) [][]int {
	if k == 0 {
		pattern := []int{}
		return [][]int{pattern}
	}

	ans := [][]int{}
	for num := begin; num < end; num++ {
		childPatterns := DupCombinationsRecursive0(num, end, k-1)
		for _, childPattern := range childPatterns {
			pattern := append([]int{num}, childPattern...)
			ans = append(ans, pattern)
		}
	}

	return ans
}

// DupCombinationsRecursive1 uses the same memory space for each combination.
func DupCombinationsRecursive1(n, k int, f func([]int)) {
	pattern := make([]int, k)

	var body func(pos, begin int)
	body = func(pos, begin int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := begin; num < n; num++ {
			pattern[pos] = num
			body(pos+1, num)
		}
	}
	body(0, 0)
}

// DupCombinationsRecursive2 does not pass the beginning of the available
// number range to recursive call. It refers the generating combination instead.
func DupCombinationsRecursive2(n, k int, f func([]int)) {
	pattern := make([]int, k+1)
	pattern[0] = 0

	var body func(pos int)
	body = func(pos int) {
		if pos == k+1 {
			f(pattern[1:])
			return
		}

		for num := pattern[pos-1]; num < n; num++ {
			pattern[pos] = num
			body(pos + 1)
		}
	}
	body(1)
}

// DupCombinationsWithCarrying0 decides the next digit of combination
// to increment not by recursive calls but by the previous combination
// directly.
func DupCombinationsWithCarrying0(n, k int, f func([]int)) {
	pattern := make([]int, k)

	for {
		f(pattern)

		pos := k - 1
		for {
			if pos == -1 {
				return
			}

			oldNum := pattern[pos]
			if oldNum == n-1 {
				// carry
				pos--
				continue
			}

			// increment
			pattern[pos]++
			break
		}

		// replace the numbers of carried digits
		numToReplace := pattern[pos]
		for pos++; pos < k; pos++ {
			pattern[pos] = numToReplace
		}
	}
}

// DupCombinationsWithCarrying1 integrates the loop for increment and carrying,
// which goes from right digit to left digit, and the one for setting rest
// values, which goes from left digit to right digit.
func DupCombinationsWithCarrying1(n, k int, f func([]int)) {
	pattern := make([]int, k)

	pos := k
	for pos > -1 {
		if pos == k {
			f(pattern)
			pos--
			continue
		}

		// carry
		oldNum := pattern[pos]
		if oldNum == n-1 {
			pattern[pos] = -1
			pos--
			continue
		}

		if oldNum == -1 {
			// replace the number of the carried digit
			pattern[pos] = pattern[pos-1]
		} else {
			// increment
			pattern[pos]++
		}
		pos++
	}
}
