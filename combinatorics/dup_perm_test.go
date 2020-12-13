package combinatorics

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDupPermutations(t *testing.T) {
	targets := []struct {
		name string
		f    func(n, k int) [][]int
	}{
		{"Recursive0",
			func(n, k int) [][]int {
				return DupPermutationsRecursive0(n, k)
			}},
		{"Recursive1",
			func(n, k int) [][]int {
				got := [][]int{}
				DupPermutationsRecursive1(n, k,
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
				DupPermutationsWithStack0(n, k,
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
				DupPermutationsWithSlice0(n, k,
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
				DupPermutationsWithCarrying0(n, k,
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
				DupPermutationsWithCarrying1(n, k,
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
		{n: 4, k: 3, want: [][]int{
			[]int{0, 0, 0},
			[]int{0, 0, 1},
			[]int{0, 0, 2},
			[]int{0, 0, 3},
			[]int{0, 1, 0},
			[]int{0, 1, 1},
			[]int{0, 1, 2},
			[]int{0, 1, 3},
			[]int{0, 2, 0},
			[]int{0, 2, 1},
			[]int{0, 2, 2},
			[]int{0, 2, 3},
			[]int{0, 3, 0},
			[]int{0, 3, 1},
			[]int{0, 3, 2},
			[]int{0, 3, 3},
			[]int{1, 0, 0},
			[]int{1, 0, 1},
			[]int{1, 0, 2},
			[]int{1, 0, 3},
			[]int{1, 1, 0},
			[]int{1, 1, 1},
			[]int{1, 1, 2},
			[]int{1, 1, 3},
			[]int{1, 2, 0},
			[]int{1, 2, 1},
			[]int{1, 2, 2},
			[]int{1, 2, 3},
			[]int{1, 3, 0},
			[]int{1, 3, 1},
			[]int{1, 3, 2},
			[]int{1, 3, 3},
			[]int{2, 0, 0},
			[]int{2, 0, 1},
			[]int{2, 0, 2},
			[]int{2, 0, 3},
			[]int{2, 1, 0},
			[]int{2, 1, 1},
			[]int{2, 1, 2},
			[]int{2, 1, 3},
			[]int{2, 2, 0},
			[]int{2, 2, 1},
			[]int{2, 2, 2},
			[]int{2, 2, 3},
			[]int{2, 3, 0},
			[]int{2, 3, 1},
			[]int{2, 3, 2},
			[]int{2, 3, 3},
			[]int{3, 0, 0},
			[]int{3, 0, 1},
			[]int{3, 0, 2},
			[]int{3, 0, 3},
			[]int{3, 1, 0},
			[]int{3, 1, 1},
			[]int{3, 1, 2},
			[]int{3, 1, 3},
			[]int{3, 2, 0},
			[]int{3, 2, 1},
			[]int{3, 2, 2},
			[]int{3, 2, 3},
			[]int{3, 3, 0},
			[]int{3, 3, 1},
			[]int{3, 3, 2},
			[]int{3, 3, 3},
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

func BenchmarkDupPermutations(b *testing.B) {
	const n = 8
	const k = 7

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
				for _, pattern := range DupPermutationsRecursive0(n, k) {
					doSomethingForPattern(pattern)
				}
			}},
		{"Recursive1",
			func() {
				DupPermutationsRecursive1(n, k, doSomethingForPattern)
			}},
		{"WithStack0",
			func() {
				DupPermutationsWithStack0(n, k, doSomethingForPattern)
			}},
		{"WithSlice0",
			func() {
				DupPermutationsWithSlice0(n, k, doSomethingForPattern)
			}},
		{"WithCarrying0",
			func() {
				DupPermutationsWithCarrying0(n, k, doSomethingForPattern)
			}},
		{"WithCarrying1",
			func() {
				DupPermutationsWithCarrying1(n, k, doSomethingForPattern)
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
