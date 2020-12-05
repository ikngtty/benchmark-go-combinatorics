package combinatorics

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	targets := []struct {
		name string
		f    func(n, k int) [][]int
	}{
		{"Recursive0",
			func(n, k int) [][]int {
				a := make([]int, n)
				for i := 0; i < n; i++ {
					a[i] = i
				}
				return PermutationsRecursive0(a, k)
			}},
		{"Recursive1",
			func(n, k int) [][]int {
				a := make([]int, n)
				for i := 0; i < n; i++ {
					a[i] = i
				}
				return PermutationsRecursive1(a, k)
			}},
		{"Recursive2",
			func(n, k int) [][]int {
				a := make([]int, n)
				for i := 0; i < n; i++ {
					a[i] = i
				}

				got := [][]int{}
				PermutationsRecursive2(a, k,
					func(pattern []int) {
						got = append(got, pattern)
					})
				return got
			}},
		{"Recursive3",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsRecursive3(n, k,
					func(pattern []int) {
						got = append(got, pattern)
					})
				return got
			}},
		{"Recursive4",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsRecursive4(n, k,
					func(pattern []int) {
						got = append(got, pattern)
					})
				return got
			}},
		{"Recursive5",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsRecursive5(n, k,
					func(pattern []int) {
						got = append(got, pattern)
					})
				return got
			}},
		{"Recursive6",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsRecursive6(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"Recursive7",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsRecursive7(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack0",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithStack0(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack1",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithStack1(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack2",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithStack2(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack3",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithStack3(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack4",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithStack4(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack5",
			func(n, k int) [][]int {
				a := make([]int, n)
				for i := 0; i < n; i++ {
					a[i] = i
				}

				got := [][]int{}
				PermutationsWithStack5(a, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack6",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithStack6(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack7",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithStack7(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithStack8",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithStack8(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithCarrying0",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithCarrying0(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithCarrying1",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithCarrying1(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithCarrying2",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithCarrying2(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithCarrying3",
			func(n, k int) [][]int {
				got := [][]int{}
				PermutationsWithCarrying3(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
	}

	cases := []struct {
		n, k int
		want [][]int
	}{
		{n: 0, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 0, want: [][]int{[]int{}}},
		{n: 3, k: 1, want: [][]int{[]int{0}, []int{1}, []int{2}}},
		{n: 5, k: 3, want: [][]int{
			[]int{0, 1, 2},
			[]int{0, 1, 3},
			[]int{0, 1, 4},
			[]int{0, 2, 1},
			[]int{0, 2, 3},
			[]int{0, 2, 4},
			[]int{0, 3, 1},
			[]int{0, 3, 2},
			[]int{0, 3, 4},
			[]int{0, 4, 1},
			[]int{0, 4, 2},
			[]int{0, 4, 3},
			[]int{1, 0, 2},
			[]int{1, 0, 3},
			[]int{1, 0, 4},
			[]int{1, 2, 0},
			[]int{1, 2, 3},
			[]int{1, 2, 4},
			[]int{1, 3, 0},
			[]int{1, 3, 2},
			[]int{1, 3, 4},
			[]int{1, 4, 0},
			[]int{1, 4, 2},
			[]int{1, 4, 3},
			[]int{2, 0, 1},
			[]int{2, 0, 3},
			[]int{2, 0, 4},
			[]int{2, 1, 0},
			[]int{2, 1, 3},
			[]int{2, 1, 4},
			[]int{2, 3, 0},
			[]int{2, 3, 1},
			[]int{2, 3, 4},
			[]int{2, 4, 0},
			[]int{2, 4, 1},
			[]int{2, 4, 3},
			[]int{3, 0, 1},
			[]int{3, 0, 2},
			[]int{3, 0, 4},
			[]int{3, 1, 0},
			[]int{3, 1, 2},
			[]int{3, 1, 4},
			[]int{3, 2, 0},
			[]int{3, 2, 1},
			[]int{3, 2, 4},
			[]int{3, 4, 0},
			[]int{3, 4, 1},
			[]int{3, 4, 2},
			[]int{4, 0, 1},
			[]int{4, 0, 2},
			[]int{4, 0, 3},
			[]int{4, 1, 0},
			[]int{4, 1, 2},
			[]int{4, 1, 3},
			[]int{4, 2, 0},
			[]int{4, 2, 1},
			[]int{4, 2, 3},
			[]int{4, 3, 0},
			[]int{4, 3, 1},
			[]int{4, 3, 2},
		}},
	}

	for _, target := range targets {
		t.Run(target.name, func(t *testing.T) {
			for _, c := range cases {
				t.Run(fmt.Sprintf("n=%d k=%d", c.n, c.k), func(t *testing.T) {
					got := target.f(c.n, c.k)
					if !reflect.DeepEqual(got, c.want) {
						t.Errorf("want: %v, got: %v", c.want, got)
					}
				})
			}
		})
	}
}

func BenchmarkPermutations(b *testing.B) {
	const n = 10
	const k = 10

	doSomethingForPattern := func(pattern []int) {
		total := 0
		for i := 1; i < len(pattern); i++ {
			total += pattern[i] - pattern[i-1]
		}
	}

	targets := []struct {
		name string
		f    func()
	}{
		{"Recursive0",
			func() {
				a := make([]int, n)
				for i := 0; i < n; i++ {
					a[i] = i
				}
				for _, pattern := range PermutationsRecursive0(a, k) {
					doSomethingForPattern(pattern)
				}
			}},
		{"Recursive1",
			func() {
				a := make([]int, n)
				for i := 0; i < n; i++ {
					a[i] = i
				}
				for _, pattern := range PermutationsRecursive1(a, k) {
					doSomethingForPattern(pattern)
				}
			}},
		{"Recursive2",
			func() {
				a := make([]int, n)
				for i := 0; i < n; i++ {
					a[i] = i
				}
				PermutationsRecursive2(a, k, doSomethingForPattern)
			}},
		{"Recursive3",
			func() {
				PermutationsRecursive3(n, k, doSomethingForPattern)
			}},
		{"Recursive4",
			func() {
				PermutationsRecursive4(n, k, doSomethingForPattern)
			}},
		{"Recursive5",
			func() {
				PermutationsRecursive5(n, k, doSomethingForPattern)
			}},
		{"Recursive6",
			func() {
				PermutationsRecursive6(n, k, doSomethingForPattern)
			}},
		{"Recursive7",
			func() {
				PermutationsRecursive7(n, k, doSomethingForPattern)
			}},
		{"WithStack0",
			func() {
				PermutationsWithStack0(n, k, doSomethingForPattern)
			}},
		{"WithStack1",
			func() {
				PermutationsWithStack1(n, k, doSomethingForPattern)
			}},
		{"WithStack2",
			func() {
				PermutationsWithStack2(n, k, doSomethingForPattern)
			}},
		{"WithStack3",
			func() {
				PermutationsWithStack3(n, k, doSomethingForPattern)
			}},
		{"WithStack4",
			func() {
				PermutationsWithStack4(n, k, doSomethingForPattern)
			}},
		{"WithStack5",
			func() {
				a := make([]int, n)
				for i := 0; i < n; i++ {
					a[i] = i
				}
				PermutationsWithStack5(a, k, doSomethingForPattern)
			}},
		{"WithStack6",
			func() {
				PermutationsWithStack6(n, k, doSomethingForPattern)
			}},
		{"WithStack7",
			func() {
				PermutationsWithStack7(n, k, doSomethingForPattern)
			}},
		{"WithStack8",
			func() {
				PermutationsWithStack8(n, k, doSomethingForPattern)
			}},
		{"WithCarrying0",
			func() {
				PermutationsWithCarrying0(n, k, doSomethingForPattern)
			}},
		{"WithCarrying1",
			func() {
				PermutationsWithCarrying1(n, k, doSomethingForPattern)
			}},
		{"WithCarrying2",
			func() {
				PermutationsWithCarrying2(n, k, doSomethingForPattern)
			}},
		{"WithCarrying3",
			func() {
				PermutationsWithCarrying3(n, k, doSomethingForPattern)
			}},
	}

	for _, target := range targets {
		b.Run(target.name, func(b *testing.B) {
			for try := 0; try < b.N; try++ {
				target.f()
			}
		})
	}
}
