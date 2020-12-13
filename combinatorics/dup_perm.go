package combinatorics

// DupPermutationsRecursive0 is a naive recursive implementation.
func DupPermutationsRecursive0(n, k int) [][]int {
	if k == 0 {
		pattern := []int{}
		return [][]int{pattern}
	}

	ans := [][]int{}
	for num := 0; num < n; num++ {
		childPatterns := DupPermutationsRecursive0(n, k-1)
		for _, childPattern := range childPatterns {
			pattern := append([]int{num}, childPattern...)
			ans = append(ans, pattern)
		}
	}

	return ans
}

// DupPermutationsRecursive1 uses the same memory space for each permutation.
func DupPermutationsRecursive1(n, k int, f func([]int)) {
	pattern := make([]int, k)

	var body func(pos int)
	body = func(pos int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := 0; num < n; num++ {
			pattern[pos] = num
			body(pos + 1)
		}
	}
	body(0)
}

// DupPermutationsWithStack0 stores nodes of permutations in a stack.
func DupPermutationsWithStack0(n, k int, f func([]int)) {
	patternNodeStack := newPatternNodeStack3()
	pattern := make([]int, k)

	patternNodeStack.push(patternNode3{pos: -1, number: 0})
	for !patternNodeStack.empty() {
		patternNode := patternNodeStack.pop()

		if patternNode.pos > -1 {
			pattern[patternNode.pos] = patternNode.number
		}

		if patternNode.pos == k-1 {
			f(pattern)
			continue
		}

		for num := n - 1; num >= 0; num-- {
			childNode := patternNode3{
				pos: patternNode.pos + 1, number: num,
			}
			patternNodeStack.push(childNode)
		}
	}
}

// DupPermutationsWithSlice0 bases on DupPermutationsWithStack0, but the stack
// is implemented with a slice, not with a pointer.
func DupPermutationsWithSlice0(n, k int, f func([]int)) {
	patternNodeStack := make([]patternNode3, 1)
	pattern := make([]int, k)

	patternNodeStack[0] = patternNode3{pos: -1, number: 0}
	for len(patternNodeStack) > 0 {
		patternNode := patternNodeStack[len(patternNodeStack)-1]      // pop(peek)
		patternNodeStack = patternNodeStack[:len(patternNodeStack)-1] // pop(discard)

		if patternNode.pos > -1 {
			pattern[patternNode.pos] = patternNode.number
		}

		if patternNode.pos == k-1 {
			f(pattern)
			continue
		}

		for num := n - 1; num >= 0; num-- {
			childNode := patternNode3{
				pos: patternNode.pos + 1, number: num,
			}
			patternNodeStack = append(patternNodeStack, childNode) // push
		}
	}
}

// DupPermutationsWithCarrying0 decides the next digit of permutation
// to increment not by recursive calls but by the previous permutation
// directly.
func DupPermutationsWithCarrying0(n, k int, f func([]int)) {
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
				pattern[pos] = 0
				pos--
				continue
			}

			// increment
			pattern[pos]++
			break
		}
	}
}

// DupPermutationsWithCarrying1 decreases its nests of loop.
func DupPermutationsWithCarrying1(n, k int, f func([]int)) {
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
			pattern[pos] = 0
			pos--
			continue
		}

		pattern[pos]++
		pos = k
	}
}

// DupPermutationsWithBaseConverting0 converts a decimal number to a base-n
// number.
func DupPermutationsWithBaseConverting0(n, k int, f func([]int)) {
	pattern := make([]int, k)

	for deciNum := 0; deciNum < Pow(n, k); deciNum++ {
		rest := deciNum
		for pos := k - 1; pos >= 0; pos-- {
			pattern[pos] = rest % n
			rest /= n
		}
		f(pattern)
	}
}
