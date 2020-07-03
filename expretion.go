package main

import (
	"log"
	"sort"
)

type Expr struct {
	Operator string
	N        int64
	Sets     [][]int64
	Child    *Expr
}

func (exp *Expr) Eval() []int64 {
	if exp.Child != nil {
		exp.Sets = append(exp.Sets, exp.Child.Eval())
	}

	return calculate(exp.Operator, exp.N, exp.Sets)
}

func calculate(op string, n int64, sets [][]int64) []int64 {
	var set []int64
	switch op {
	case OpGr:
		set = gr(n, sets)
	case OpLe:
		set = le(n, sets)
	case OpEq:
		set = eq(n, sets)
	default:
		log.Fatal("unknown operation: " + op)
	}

	sort.Sort(Int64Slice(set))
	return set
}

func calcIndex(sets [][]int64) map[int64]int64 {
	index := map[int64]int64{}

	for _, set := range sets {
		for _, val := range set {
			index[val] += 1
		}
	}

	return index
}

func gr(n int64, sets [][]int64) []int64 {
	var set []int64
	for val, count := range calcIndex(sets) {
		if count > n {
			set = append(set, val)
		}
	}

	return set
}

func le(n int64, sets [][]int64) []int64 {
	var set []int64
	for val, count := range calcIndex(sets) {
		if count < n {
			set = append(set, val)
		}
	}

	return set
}

func eq(n int64, sets [][]int64) []int64 {
	var set []int64
	for val, count := range calcIndex(sets) {
		if count == n {
			set = append(set, val)
		}
	}

	return set
}
