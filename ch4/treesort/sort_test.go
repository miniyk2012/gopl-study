// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"gopl.io/ch4/treesort"
)

func appendValues(values []int) []int {
	fmt.Printf("%p\n", &values)
	fmt.Printf("%d %d\n", len(values), cap(values))
	values[0] = 10
	values[1] = 20
	return values
}

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	var root *treesort.Tree
	fmt.Printf("%t\n", root == nil)
	root = new(treesort.Tree)
	fmt.Printf("%v\n", *root)
	fmt.Printf("%v\n", data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}

	values := make([]int, 5, 10)
	fmt.Printf("%p\n", &values)
	fmt.Printf("%v\n", values)
	appendValues(values)
	fmt.Printf("%v\n", values)
}
