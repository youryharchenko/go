package main

import (
	"fmt"
)

type enum int

const (
	first enum = iota
	second
)

type pair func(m enum) interface{}

func (p pair) String() string {
	tail := cdr(p)
	//fmt.Println("String pair", reflect.TypeOf(tail), car(p), tail)
	if tail == nil {
		return fmt.Sprintf("{%v,nil}", car(p))
	}
	return fmt.Sprintf("{%v %v}", car(p), cdr(p))
}

func cons(x, y interface{}) pair {
	return func(m enum) interface{} {
		if m == first {
			return x
		}
		if m == second {
			return y
		}
		panic(fmt.Sprintf("Argument not 0 or 1 - CONS: %d", m))
	}
}

func car(p pair) interface{} {
	return p(first)
}

func cdr(p pair) interface{} {
	return p(second)
}

type list func() pair

func (l list) String() string {
	//fmt.Println("String list", car(pair(l)))
	return fmt.Sprintf("%v", listToSlice(l))
}

func newList(p pair) list {
	return func() pair {
		return p
	}
}

func makeList(e ...interface{}) list {
	return sliceToList(e)
}

func listRef(items list, n int) interface{} {
	if n == 0 {
		return car(items())
	}
	tail := cdr(items())
	if tail == nil {
		panic("out of range")
	}
	return listRef(newList(tail.(pair)), n-1)
}

func length(items list) int {
	if items() == nil {
		return 0
	}
	tail := cdr(items())
	if tail == nil {
		return 1
	}
	return 1 + length(newList(tail.(pair)))
}

func join(list1, list2 list) list {
	if list1() == nil {
		return list2
	}
	tail := cdr(list1())
	if tail == nil {
		return newList(cons(car(list1()), list2()))
	}
	return newList(cons(car(list1()), join(newList(tail.(pair)), list2)()))
}

func listToSlice(l list) []interface{} {
	a := []interface{}{}
	if l() == nil {
		return a
	}
	a = append(a, car(l()))
	tail := cdr(l())
	if tail == nil {
		return a
	}
	return append(a, listToSlice(newList(tail.(pair)))...)
}

func sliceToList(s []interface{}) list {
	//fmt.Println(s)
	if s == nil || len(s) == 0 {
		return newList(nil)
	}
	if len(s) == 1 {
		return newList(cons(s[0], nil))
	}
	return newList(cons(s[0], sliceToList(s[1:])()))
}
