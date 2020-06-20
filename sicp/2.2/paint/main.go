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

func fmap(proc func(interface{}) interface{}, items list) list {
	if items() == nil {
		return newList(nil)
	}
	tail := cdr(items())
	if tail == nil {
		return newList(cons(proc(car(items())), nil))
	}
	return newList(cons(proc(car(items())), fmap(proc, newList(tail.(pair)))()))
}

func feach(proc func(interface{}), items list) {
	if items() == nil {
		return
	}
	proc(car(items()))
	tail := cdr(items())
	if tail != nil {
		feach(proc, newList(tail.(pair)))
	}
}

func countLeaves(items list) int {
	if items() == nil {
		return 0
	}
	head := car(items())
	tail := cdr(items())
	l, ok := head.(list)
	if tail == nil {
		if ok {
			return countLeaves(l)
		}
		return 1
	}
	if ok {
		return countLeaves(l) + countLeaves(newList(tail.(pair)))
	}
	return 1 + countLeaves(newList(tail.(pair)))
}

func filter(predicate func(x interface{}) bool, sequence list) list {
	if sequence() == nil {
		return newList(nil)
	}
	head := car(sequence())
	tail := cdr(sequence())
	if tail == nil {
		if predicate(head) {
			return newList(cons(head, nil))
		}
		return newList(nil)
	}
	if predicate(head) {
		return newList(cons(head, filter(predicate, newList(tail.(pair)))()))
	}
	return filter(predicate, newList(tail.(pair)))
}

func foldRight(op func(x, y interface{}) interface{}, initial interface{}, sequence list) interface{} {
	if sequence() == nil {
		return initial
	}
	head := car(sequence())
	tail := cdr(sequence())
	if tail == nil {
		return op(head, initial)
	}
	return op(head, foldRight(op, initial, newList(tail.(pair))))
}

func foldLeft(op func(x, y interface{}) interface{}, initial interface{}, sequence list) interface{} {
	var iter func(interface{}, list) interface{}
	iter = func(result interface{}, rest list) interface{} {
		if rest() == nil {
			return initial
		}
		head := car(rest())
		tail := cdr(rest())
		if tail == nil {
			return op(result, head)
		}
		return iter(op(result, head), newList(tail.(pair)))
	}
	return iter(initial, sequence)
}

func flatMap(proc func(interface{}) interface{}, sequence list) list {
	return foldRight(
		func(x, y interface{}) interface{} {
			if x == nil {
				x = newList(nil)
			}
			if y == nil {
				y = newList(nil)
			}
			return join(x.(list), y.(list))
		},
		nil,
		fmap(proc, sequence),
	).(list)
}

func accumulate(op func(x, y interface{}) interface{}, initial interface{}, sequence list) interface{} {
	return foldRight(op, initial, sequence)
}

func enumerateInterval(low, high int) list {
	if low > high {
		return newList(nil)
	}
	return newList(cons(low, enumerateInterval(low+1, high)()))
}

func enumerateTree(tree list) list {
	if tree() == nil {
		return newList(nil)
	}
	head := car(tree())
	tail := cdr(tree())
	l, ok := head.(list)
	if tail == nil {
		if ok {
			return enumerateTree(l)
		}
		return newList(cons(head, nil))
	}
	if ok {
		return join(enumerateTree(l), enumerateTree(newList(tail.(pair))))
	}
	return newList(cons(head, enumerateTree(newList(tail.(pair)))()))
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
