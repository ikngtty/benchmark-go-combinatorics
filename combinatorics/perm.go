package combinatorics

// PermutationsRecursive0 is a naive recursive implementation.
func PermutationsRecursive0(a []int, k int) [][]int {
	if k == 0 {
		pattern := []int{}
		return [][]int{pattern}
	}

	ans := [][]int{}
	for i := range a {
		// make a new array without i-th item
		aRest := make([]int, len(a)-1)
		for j := 0; j < i; j++ {
			aRest[j] = a[j]
		}
		for j := i + 1; j < len(a); j++ {
			aRest[j-1] = a[j]
		}

		childPatterns := PermutationsRecursive0(aRest, k-1)
		for _, childPattern := range childPatterns {
			pattern := append([]int{a[i]}, childPattern...)
			ans = append(ans, pattern)
		}
	}

	return ans
}

// PermutationsRecursive1 computes the number of permutations and allocates
// spaces statically for the array of permutations.
func PermutationsRecursive1(a []int, k int) [][]int {
	if k == 0 {
		pattern := []int{}
		return [][]int{pattern}
	}

	ans := make([][]int, 0, PermutationCount(len(a), k))
	for i := range a {
		// make a new array without i-th item
		aRest := make([]int, len(a)-1)
		for j := 0; j < i; j++ {
			aRest[j] = a[j]
		}
		for j := i + 1; j < len(a); j++ {
			aRest[j-1] = a[j]
		}

		childPatterns := PermutationsRecursive1(aRest, k-1)
		for _, childPattern := range childPatterns {
			pattern := append([]int{a[i]}, childPattern...)
			ans = append(ans, pattern)
		}
	}

	return ans
}

// PermutationsRecursive2 passes each permutation to the callback function
// instead of returning an array of all permutations. In other words, it is
// a lazy evaluation for permutations.
func PermutationsRecursive2(a []int, k int, f func([]int)) {
	if k == 0 {
		pattern := []int{}
		f(pattern)
		return
	}

	for i := range a {
		// make a new array without i-th item
		aRest := make([]int, len(a)-1)
		for j := 0; j < i; j++ {
			aRest[j] = a[j]
		}
		for j := i + 1; j < len(a); j++ {
			aRest[j-1] = a[j]
		}

		PermutationsRecursive2(aRest, k-1, func(childPattern []int) {
			pattern := append([]int{a[i]}, childPattern...)
			f(pattern)
		})
	}
}

// PermutationsRecursive3 uses an array of bool rather than an array of int
// as available numbers. e.g. If `n = 5` and available numbers are `[1, 4]`,
// it treats `[false, true, false, false, true]`.
func PermutationsRecursive3(n, k int, f func([]int)) {
	checklist := make([]bool, n)

	var body func(k int, f func([]int))
	body = func(k int, f func([]int)) {
		if k == 0 {
			f([]int{})
			return
		}

		for num := range checklist {
			if checklist[num] {
				continue
			}

			checklist[num] = true
			body(k-1, func(childPattern []int) {
				pattern := append([]int{num}, childPattern...)
				f(pattern)
			})
			checklist[num] = false
		}
	}
	body(k, f)
}

// PermutationsRecursive4 uses a list rather than an array for each permutation.
func PermutationsRecursive4(n, k int, f func([]int)) {
	checklist := make([]bool, n)

	var body func(k int, f func(*IntList))
	body = func(k int, f func(*IntList)) {
		if k == 0 {
			f(NewIntList())
			return
		}

		for num := range checklist {
			if checklist[num] {
				continue
			}

			checklist[num] = true
			body(k-1, func(childPattern *IntList) {
				pattern := NewIntList()
				pattern.Add(num)
				pattern.Concat(childPattern)
				f(pattern)
			})
			checklist[num] = false
		}
	}
	body(k, func(list *IntList) {
		f(list.ToA())
	})
}

// PermutationsRecursive5 allocates enough spaces first for an array of
// a permutation even though in the recursive call.
func PermutationsRecursive5(n, k int, f func([]int)) {
	checklist := make([]bool, n)

	var body func(pos int, f func([]int))
	body = func(pos int, f func([]int)) {
		if pos == k {
			f(make([]int, k))
			return
		}

		for num := range checklist {
			if checklist[num] {
				continue
			}

			checklist[num] = true
			body(pos+1, func(pattern []int) {
				pattern[pos] = num
				f(pattern)
			})
			checklist[num] = false
		}
	}
	body(0, f)
}

// PermutationsRecursive6 uses the same memory space for each permutation.
func PermutationsRecursive6(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	pattern := make([]int, k)

	var body func(pos int)
	body = func(pos int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := range checklist {
			if checklist[num] {
				continue
			}

			pattern[pos] = num
			checklist[num] = true
			body(pos + 1)
			checklist[num] = false
		}
	}
	body(0)
}

// PermutationsRecursive7 does not record available numbers. Instead, not to
// use same numbers in a permutation, it checks each digit every time.
func PermutationsRecursive7(n, k int, f func([]int)) {
	pattern := make([]int, k)

	var body func(pos int)
	body = func(pos int) {
		if pos == k {
			f(pattern)
			return
		}

		for num := 0; num < n; num++ {
			// skip if the number of `num` is used in the left digits
			willContinue := false
			for i := 0; i < pos; i++ {
				if pattern[i] == num {
					willContinue = true
					break
				}
			}
			if willContinue {
				continue
			}

			pattern[pos] = num
			body(pos + 1)
		}
	}
	body(0)
}

// PermutationsWithStack0 does not use recursive calls. Instead, it uses
// a stack that imitates the recursive call stack.
func PermutationsWithStack0(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	callStack := newCallStack0()
	pattern := make([]int, k)

	callStack.push(&callStackItem0{pos: 0, chosenNumber: -1})
	for !callStack.empty() {
		env := callStack.peek()

		// at the most right digit, call back the function
		if env.pos == k {
			f(pattern)

			callStack.pop()
			continue
		}

		// reset the digit of `checklist` before increment the digit
		if env.chosenNumber > -1 {
			checklist[env.chosenNumber] = false
		}

		// increment the digit
		willContinue := false
		for env.chosenNumber++; env.chosenNumber < n; env.chosenNumber++ {
			// skip if the number of `env.chosenNumber` is used
			if checklist[env.chosenNumber] {
				continue
			}

			// fill the number
			pattern[env.pos] = env.chosenNumber
			checklist[env.chosenNumber] = true

			// push a stack item for the right digit
			newEnv := callStackItem0{
				pos: env.pos + 1, chosenNumber: -1,
			}
			callStack.push(&newEnv)
			willContinue = true
			break
		}
		if willContinue {
			continue
		}

		// the case it cannot increment the digit
		// -> remove the stack item
		callStack.pop()
	}
}

// PermutationsWithStack1 does not store the chosen number in the psuedo
// "call stack". It refers the current generating permutation instead.
func PermutationsWithStack1(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	posStack := NewIntStack()
	pattern := make([]int, k)
	for i := range pattern {
		pattern[i] = -1
	}

	posStack.Push(0)
	for !posStack.Empty() {
		pos := posStack.Peek()

		// at the most right digit, call back the function
		if pos == k {
			f(pattern)

			posStack.Pop()
			continue
		}

		// reset the digit of `checklist` before increment the digit
		chosenNumber := pattern[pos]
		if chosenNumber > -1 {
			checklist[chosenNumber] = false
		}

		// increment the digit
		willContinue := false
		for chosenNumber++; chosenNumber < n; chosenNumber++ {
			// skip if the number of `chosenNumber` is used
			if checklist[chosenNumber] {
				continue
			}

			// fill the number
			pattern[pos] = chosenNumber
			checklist[chosenNumber] = true

			// push a stack item for the right digit
			posStack.Push(pos + 1)
			willContinue = true
			break
		}
		if willContinue {
			continue
		}

		// the case it cannot increment the digit
		// -> reset the digit of `pattern` and remove the stack item
		pattern[pos] = -1
		posStack.Pop()
	}
}

// PermutationsWithStack2 stores nodes of permutations in a stack.
func PermutationsWithStack2(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	patternNodeStack := newPatternNodeStack2()
	pattern := make([]int, k)

	patternNodeStack.push(&patternNode2{pos: -1, number: 0})
	for !patternNodeStack.empty() {
		patternNode := patternNodeStack.pop()

		// reset the right digits of `checklist` and `pattern`
		for i := patternNode.pos; i < k; i++ {
			if i > -1 && pattern[i] > -1 {
				checklist[pattern[i]] = false
				pattern[i] = -1
			}
		}

		// fill the number
		if patternNode.pos > -1 {
			pattern[patternNode.pos] = patternNode.number
			checklist[patternNode.number] = true
		}

		// at the most right digit, call back the function
		if patternNode.pos == k-1 {
			f(pattern)
			continue
		}

		// enumerate the numbers of the right digit
		for num := n - 1; num >= 0; num-- {
			// skip if the number of `num` is used
			if checklist[num] {
				continue
			}

			// push a stack item for the right digit
			childNode := patternNode2{
				pos: patternNode.pos + 1, number: num,
			}
			patternNodeStack.push(&childNode)
		}
	}
}

// PermutationsWithStack3 treats a stack item as a value, not as a reference.
func PermutationsWithStack3(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	patternNodeStack := newPatternNodeStack3()
	pattern := make([]int, k)

	patternNodeStack.push(patternNode3{pos: -1, number: 0})
	for !patternNodeStack.empty() {
		patternNode := patternNodeStack.pop()

		// reset the right digits of `checklist` and `pattern`
		for i := patternNode.pos; i < k; i++ {
			if i > -1 && pattern[i] > -1 {
				checklist[pattern[i]] = false
				pattern[i] = -1
			}
		}

		// fill the number
		if patternNode.pos > -1 {
			pattern[patternNode.pos] = patternNode.number
			checklist[patternNode.number] = true
		}

		// at the most right digit, call back the function
		if patternNode.pos == k-1 {
			f(pattern)
			continue
		}

		// enumerate the numbers of the right digit
		for num := n - 1; num >= 0; num-- {
			// skip if the number of `num` is used
			if checklist[num] {
				continue
			}

			// push a stack item for the right digit
			childNode := patternNode3{
				pos: patternNode.pos + 1, number: num,
			}
			patternNodeStack.push(childNode)
		}
	}
}

// PermutationsWithStack4 searches from nodes for the first digit of
// permutations, while other functions do from the sentinel node, so to speak,
// which is for the "zeroth digit".
func PermutationsWithStack4(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	patternNodeStack := newPatternNodeStack3()
	pattern := make([]int, k)

	if k == 0 {
		f(pattern)
		return
	}

	for num := n - 1; num >= 0; num-- {
		patternNodeStack.push(patternNode3{pos: 0, number: num})
	}

	for !patternNodeStack.empty() {
		patternNode := patternNodeStack.pop()

		// reset the right digits of `checklist` and `pattern`
		for i := patternNode.pos; i < k; i++ {
			if pattern[i] > -1 {
				checklist[pattern[i]] = false
				pattern[i] = -1
			}
		}

		// fill the number
		pattern[patternNode.pos] = patternNode.number
		checklist[patternNode.number] = true

		// at the most right digit, call back the function
		if patternNode.pos == k-1 {
			f(pattern)
			continue
		}

		// enumerate the numbers of the right digit
		for num := n - 1; num >= 0; num-- {
			// skip if the number of `num` is used
			if checklist[num] {
				continue
			}

			// push a stack item for the right digit
			childNode := patternNode3{
				pos: patternNode.pos + 1, number: num,
			}
			patternNodeStack.push(childNode)
		}
	}
}

// PermutationsWithStack5 uses an array of int rather than an array of bool
// as available numbers.
func PermutationsWithStack5(a []int, k int, f func([]int)) {
	patternNodeStack := newPatternNodeStack5()
	pattern := make([]int, k)

	// NOTE: By measuring performance, it is found that treating a stack item
	// as a value is still faster than treating it as a reference.
	patternNodeStack.push(patternNode5{pos: -1, number: 0, rest: a})
	for !patternNodeStack.empty() {
		patternNode := patternNodeStack.pop()

		// fill the number
		if patternNode.pos > -1 {
			pattern[patternNode.pos] = patternNode.number
		}

		// at the most right digit, call back the function
		if patternNode.pos == k-1 {
			f(pattern)
			continue
		}

		// enumerate the numbers of the right digit
		for i := len(patternNode.rest) - 1; i >= 0; i-- {
			// make a new array without i-th item
			newRest := make([]int, len(patternNode.rest)-1)
			for j := 0; j < i; j++ {
				newRest[j] = patternNode.rest[j]
			}
			for j := i + 1; j < len(patternNode.rest); j++ {
				newRest[j-1] = patternNode.rest[j]
			}

			// push a stack item for the right digit
			childNode := patternNode5{
				pos:    patternNode.pos + 1,
				number: patternNode.rest[i],
				rest:   newRest,
			}
			patternNodeStack.push(childNode)
		}
	}
}

// PermutationsWithStack6 does not record available numbers. Instead, not to
// use same numbers in a permutation, it checks each digit every time.
func PermutationsWithStack6(n, k int, f func([]int)) {
	patternNodeStack := newPatternNodeStack3()
	pattern := make([]int, k)

	patternNodeStack.push(patternNode3{pos: -1, number: 0})
	for !patternNodeStack.empty() {
		patternNode := patternNodeStack.pop()

		// fill the number
		if patternNode.pos > -1 {
			pattern[patternNode.pos] = patternNode.number
		}

		// at the most right digit, call back the function
		if patternNode.pos == k-1 {
			f(pattern)
			continue
		}

		// enumerate the numbers of the right digit
		for num := n - 1; num >= 0; num-- {
			// skip if the number of `num` is used in the left digits
			willContinue := false
			for i := 0; i <= patternNode.pos; i++ {
				if pattern[i] == num {
					willContinue = true
					break
				}
			}
			if willContinue {
				continue
			}

			// push a stack item for the right digit
			childNode := patternNode3{
				pos: patternNode.pos + 1, number: num,
			}
			patternNodeStack.push(childNode)
		}
	}
}

// PermutationsWithStack7 stores denotions of operation in a stack.
func PermutationsWithStack7(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	operationStack := newOperationStack7()
	pattern := make([]int, k)

	operationStack.push(operation7{
		pos:    -1,
		number: 0,
		mode:   operationMode7ExecuteOrDelegate,
	})
	for !operationStack.empty() {
		operation := operationStack.pop()

		switch operation.mode {
		case operationMode7ReflectValue:
			// fill the number
			pattern[operation.pos] = operation.number
			checklist[operation.number] = true

		case operationMode7ExecuteOrDelegate:
			if operation.pos == k-1 {
				// at the most right digit, call back the function
				f(pattern)
			} else {
				// enumerate the numbers of the right digit
				for num := n - 1; num >= 0; num-- {
					// skip if the number of `num` is used
					if checklist[num] {
						continue
					}

					// push stack items for the right digit
					operationStack.push(operation7{
						pos:    operation.pos + 1,
						number: num,
						mode:   operationMode7ResetValue,
					})
					operationStack.push(operation7{
						pos:    operation.pos + 1,
						number: num,
						mode:   operationMode7ExecuteOrDelegate,
					})
					operationStack.push(operation7{
						pos:    operation.pos + 1,
						number: num,
						mode:   operationMode7ReflectValue,
					})
				}
			}

		case operationMode7ResetValue:
			// reset the digit of `checklist`
			checklist[operation.number] = false
		}
	}
}

// PermutationsWithStack8 stores functions in a stack.
func PermutationsWithStack8(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	operationStack := NewFuncStack()
	pattern := make([]int, k)

	reflectValue := func(pos, number int) {
		// fill the number
		pattern[pos] = number
		checklist[number] = true
	}
	resetValue := func(number int) {
		// reset the digit of `checklist`
		checklist[number] = false
	}
	var executeOrDelegate func(pos int)
	executeOrDelegate = func(pos int) {
		// at the most right digit, call back the function
		if pos == k-1 {
			f(pattern)
			return
		}

		// enumerate the numbers of the right digit
		for num := n - 1; num >= 0; num-- {
			// skip if the number of `num` is used
			if checklist[num] {
				continue
			}

			// push stack items for the right digit
			numFrozen := num
			operationStack.Push(func() {
				resetValue(numFrozen)
			})
			operationStack.Push(func() {
				executeOrDelegate(pos + 1)
			})
			operationStack.Push(func() {
				reflectValue(pos+1, numFrozen)
			})
		}
	}

	operationStack.Push(func() {
		executeOrDelegate(-1)
	})
	for !operationStack.Empty() {
		operation := operationStack.Pop()
		operation()
	}
}

// PermutationsWithCarrying0 decides the next digit of permutation
// to increment not by a stack or recursive calls but by the previous
// permutation directly.
func PermutationsWithCarrying0(n, k int, f func([]int)) {
	pattern := make([]int, k)
	for i := range pattern {
		pattern[i] = i
	}

	for {
		f(pattern)

		// increment
		pos := k - 1 // current digit
		for pos >= 0 {
			oldNum := pattern[pos]

			willBreak := false
			for newNum := oldNum + 1; newNum < n; newNum++ {
				// skip if the number of `newNum` is used in the left digits
				willContinue := false
				for i := 0; i < pos; i++ {
					if pattern[i] == newNum {
						willContinue = true
						break
					}
				}
				if willContinue {
					continue
				}

				// increment the value of the current digit
				pattern[pos] = newNum
				willBreak = true
				break
			}
			if willBreak {
				break
			}

			// the case it cannot increment the current digit
			// -> carry
			pos--
		}
		// end of enumerating permutations
		if pos == -1 {
			break
		}

		// replace the numbers of carried digits
		for pos++; pos < k; pos++ {
			for num := 0; num < k; num++ {
				// skip if the number of `num` is used in the left digits
				willContinue := false
				for i := 0; i < pos; i++ {
					if pattern[i] == num {
						willContinue = true
						break
					}
				}
				if willContinue {
					continue
				}

				// replace
				pattern[pos] = num
				break
			}
		}
	}
}

// PermutationsWithCarrying1 records available numbers to an array of bool.
func PermutationsWithCarrying1(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	pattern := make([]int, k)
	for i := range pattern {
		pattern[i] = i
		checklist[i] = true
	}

	for {
		f(pattern)

		// increment
		pos := k - 1 // current digit
		for pos >= 0 {
			oldNum := pattern[pos]
			checklist[oldNum] = false

			willBreak := false
			for newNum := oldNum + 1; newNum < n; newNum++ {
				// skip if the number of `newNum` is used
				if checklist[newNum] {
					continue
				}

				// increment the value of the current digit
				pattern[pos] = newNum
				checklist[newNum] = true
				willBreak = true
				break
			}
			if willBreak {
				break
			}

			// the case it cannot increment the current digit
			// -> carry
			pos--
		}
		// end of enumerating permutations
		if pos == -1 {
			break
		}

		// replace the numbers of carried digits
		for pos++; pos < k; pos++ {
			for num := 0; num < k; num++ {
				// skip if the number of `num` is used
				if checklist[num] {
					continue
				}

				// replace
				pattern[pos] = num
				checklist[num] = true
				break
			}
		}
	}
}

// PermutationsWithCarrying2 integrates the loop for increment and carrying,
// which goes from right digit to left digit, and the one for setting rest
// values, which goes from left digit to right digit.
func PermutationsWithCarrying2(n, k int, f func([]int)) {
	checklist := make([]bool, n)
	pattern := make([]int, k)
	for i := range pattern {
		pattern[i] = i
		checklist[i] = true
	}

	pos := k
	for pos > -1 {
		if pos == k {
			f(pattern)
			pos--
			continue
		}

		oldNum := pattern[pos]
		if oldNum > -1 {
			checklist[oldNum] = false
		}

		willContinue := false
		for newNum := oldNum + 1; newNum < n; newNum++ {
			if checklist[newNum] {
				continue
			}

			pattern[pos] = newNum
			checklist[newNum] = true
			pos++
			willContinue = true
			break
		}
		if willContinue {
			continue
		}

		// carry
		pattern[pos] = -1
		pos--
	}
}

// PermutationCount computes the number of permutations.
func PermutationCount(n, k int) int {
	ans := 1
	for i := 0; i < k; i++ {
		ans *= n - i
	}
	return ans
}

type callStackItem0 struct {
	pos          int
	chosenNumber int
}

type callStack0 struct {
	last *callStack0Node
}

type callStack0Node struct {
	parent *callStack0Node
	value  *callStackItem0
}

func newCallStack0() *callStack0 {
	return &callStack0{nil}
}

func (s *callStack0) push(elem *callStackItem0) {
	node := callStack0Node{s.last, elem}
	s.last = &node
}

func (s *callStack0) pop() *callStackItem0 {
	value := s.last.value
	s.last = s.last.parent
	return value
}

func (s *callStack0) peek() *callStackItem0 {
	return s.last.value
}

func (s *callStack0) empty() bool {
	return s.last == nil
}

type patternNode2 struct {
	pos    int
	number int
}

type patternNodeStack2 struct {
	last *patternNodeStack2Node
}

type patternNodeStack2Node struct {
	parent *patternNodeStack2Node
	value  *patternNode2
}

func newPatternNodeStack2() *patternNodeStack2 {
	return &patternNodeStack2{nil}
}

func (s *patternNodeStack2) push(elem *patternNode2) {
	node := patternNodeStack2Node{s.last, elem}
	s.last = &node
}

func (s *patternNodeStack2) pop() *patternNode2 {
	value := s.last.value
	s.last = s.last.parent
	return value
}

func (s *patternNodeStack2) empty() bool {
	return s.last == nil
}

type patternNode3 struct {
	pos    int
	number int
}

type patternNodeStack3 struct {
	last *patternNodeStack3Node
}

type patternNodeStack3Node struct {
	parent *patternNodeStack3Node
	value  patternNode3
}

func newPatternNodeStack3() *patternNodeStack3 {
	return &patternNodeStack3{nil}
}

func (s *patternNodeStack3) push(elem patternNode3) {
	node := patternNodeStack3Node{s.last, elem}
	s.last = &node
}

func (s *patternNodeStack3) pop() patternNode3 {
	value := s.last.value
	s.last = s.last.parent
	return value
}

func (s *patternNodeStack3) empty() bool {
	return s.last == nil
}

type patternNode5 struct {
	pos    int
	number int
	rest   []int
}

type patternNodeStack5 struct {
	last *patternNodeStack5Node
}

type patternNodeStack5Node struct {
	parent *patternNodeStack5Node
	value  patternNode5
}

func newPatternNodeStack5() *patternNodeStack5 {
	return &patternNodeStack5{nil}
}

func (s *patternNodeStack5) push(elem patternNode5) {
	node := patternNodeStack5Node{s.last, elem}
	s.last = &node
}

func (s *patternNodeStack5) pop() patternNode5 {
	value := s.last.value
	s.last = s.last.parent
	return value
}

func (s *patternNodeStack5) empty() bool {
	return s.last == nil
}

const (
	operationMode7ReflectValue = iota
	operationMode7ExecuteOrDelegate
	operationMode7ResetValue
)

type operation7 struct {
	pos    int
	number int
	mode   int
}

type operationStack7 struct {
	last *operationStack7Node
}

type operationStack7Node struct {
	parent *operationStack7Node
	value  operation7
}

func newOperationStack7() *operationStack7 {
	return &operationStack7{nil}
}

func (s *operationStack7) push(elem operation7) {
	node := operationStack7Node{s.last, elem}
	s.last = &node
}

func (s *operationStack7) pop() operation7 {
	value := s.last.value
	s.last = s.last.parent
	return value
}

func (s *operationStack7) empty() bool {
	return s.last == nil
}
