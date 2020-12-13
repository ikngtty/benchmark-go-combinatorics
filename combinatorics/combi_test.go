package combinatorics

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	targets := []struct {
		name string
		f    func(n, k int) [][]int
	}{
		{"Recursive0",
			func(n, k int) [][]int {
				return CombinationsRecursive0(0, n, k)
			}},
		{"Recursive1",
			func(n, k int) [][]int {
				got := [][]int{}
				CombinationsRecursive1(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"Recursive2",
			func(n, k int) [][]int {
				got := [][]int{}
				CombinationsRecursive2(n, k,
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
				CombinationsWithStack0(n, k,
					func(pattern []int) {
						patternClone := make([]int, len(pattern))
						copy(patternClone, pattern)
						got = append(got, patternClone)
					})
				return got
			}},
		{"WithSlice0",
			func(n, k int) [][]int {
				got := [][]int{}
				CombinationsWithSlice0(n, k,
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
				CombinationsWithCarrying0(n, k,
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
				CombinationsWithCarrying1(n, k,
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
		{n: 3, k: 3, want: [][]int{[]int{0, 1, 2}}},
		{n: 6, k: 3, want: [][]int{
			[]int{0, 1, 2},
			[]int{0, 1, 3},
			[]int{0, 1, 4},
			[]int{0, 1, 5},
			[]int{0, 2, 3},
			[]int{0, 2, 4},
			[]int{0, 2, 5},
			[]int{0, 3, 4},
			[]int{0, 3, 5},
			[]int{0, 4, 5},
			[]int{1, 2, 3},
			[]int{1, 2, 4},
			[]int{1, 2, 5},
			[]int{1, 3, 4},
			[]int{1, 3, 5},
			[]int{1, 4, 5},
			[]int{2, 3, 4},
			[]int{2, 3, 5},
			[]int{2, 4, 5},
			[]int{3, 4, 5},
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

func BenchmarkCombinations(b *testing.B) {
	const n = 24
	const k = 12

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
				for _, pattern := range CombinationsRecursive0(0, n, k) {
					doSomethingForPattern(pattern)
				}
			}},
		{"Recursive1",
			func() {
				CombinationsRecursive1(n, k, doSomethingForPattern)
			}},
		{"Recursive2",
			func() {
				CombinationsRecursive2(n, k, doSomethingForPattern)
			}},
		{"WithStack0",
			func() {
				CombinationsWithStack0(n, k, doSomethingForPattern)
			}},
		{"WithSlice0",
			func() {
				CombinationsWithSlice0(n, k, doSomethingForPattern)
			}},
		{"WithCarrying0",
			func() {
				CombinationsWithCarrying0(n, k, doSomethingForPattern)
			}},
		{"WithCarrying1",
			func() {
				CombinationsWithCarrying1(n, k, doSomethingForPattern)
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
