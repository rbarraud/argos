package main

import (
	"fmt"
	"sort"
	"strconv"
)

func printSortedMapStringInt(m map[string]int, threshold int) map[int][]string {
	total := len(m)
	fmt.Println("total: " + strconv.Itoa(total))
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		if v > threshold {
			n[v] = append(n[v], k)
		}
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Println(strconv.Itoa(k) + " - " + s)
		}
	}
	return n
}
