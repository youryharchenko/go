package main

import (
	"fmt"
	"testing"
)

var oneThroughFour = makeList(1, 2, 3, 4)

func TestPair1(t *testing.T) {
	fmt.Println(cons(1, 2))
	t.Error("")
}
func TestPair2(t *testing.T) {
	fmt.Println(cons(1, "abc"))
	t.Error("")
}

func TestPair3(t *testing.T) {
	fmt.Println(cons(1, cons(2, nil)))
	t.Error("")
}

func TestList1(t *testing.T) {
	fmt.Println(newList(cons(1, cons(2, nil))))
	t.Error("")
}

func TestList2(t *testing.T) {
	fmt.Println(newList(cons(1, cons(2, (cons(3, nil))))))
	t.Error("")
}
func TestMakeList1(t *testing.T) {
	fmt.Println(makeList(1, 2, 3, 4, 5))
	t.Error("")
}

func TestMakeList2(t *testing.T) {
	fmt.Println(makeList())
	t.Error("")
}

func TestMakeList3(t *testing.T) {
	fmt.Println(makeList(1, true, "abc", false, 10.5, '\n'))
	t.Error("")
}
func TestCar(t *testing.T) {
	fmt.Println(car(oneThroughFour()))
	t.Error("")
}

func TestCdr(t *testing.T) {
	fmt.Println(cdr(oneThroughFour()))
	t.Error("")
}

func TestCdr2(t *testing.T) {
	fmt.Println(newList(cdr(oneThroughFour()).(pair)))
	t.Error("")
}

func TestCons1(t *testing.T) {
	fmt.Println(cons(10, oneThroughFour()))
	t.Error("")
}

func TestCons2(t *testing.T) {
	fmt.Println(newList(cons(10, oneThroughFour())))
	t.Error("")
}

func TestListRef1(t *testing.T) {
	fmt.Println(listRef(oneThroughFour, 0))
	t.Error("")
}

func TestListRef2(t *testing.T) {
	fmt.Println(listRef(oneThroughFour, 3))
	t.Error("")
}

func TestListRef3(t *testing.T) {
	fmt.Println(listRef(oneThroughFour, 4))
	t.Error("")
}

func TestListLength1(t *testing.T) {
	fmt.Println(length(makeList(1, 2, 3, 4, 5)))
	t.Error("")
}

func TestListLength2(t *testing.T) {
	fmt.Println(length(makeList()))
	t.Error("")
}
func TestListLength3(t *testing.T) {
	fmt.Println(length(makeList(1)))
	t.Error("")
}

func TestListJoin1(t *testing.T) {
	fmt.Println(join(makeList(1, 2, 3, 4, 5), makeList(6, 7, 8)))
	t.Error("")
}

func TestListJoin2(t *testing.T) {
	fmt.Println(join(makeList(), makeList(6, 7, 8)))
	t.Error("")
}

func TestListJoin3(t *testing.T) {
	fmt.Println(join(makeList(1, 2, 3, 4, 5), makeList()))
	t.Error("")
}

func TestListJoin4(t *testing.T) {
	fmt.Println(join(makeList(1), makeList(8)))
	t.Error("")
}

func TestListJoin5(t *testing.T) {
	fmt.Println(join(makeList(), makeList()))
	t.Error("")
}

func TestListMap1(t *testing.T) {
	fmt.Println(
		fmap(
			func(x interface{}) interface{} {
				return x.(int) + 1
			},
			makeList(),
		),
	)
	t.Error("")
}

func TestListMap2(t *testing.T) {
	fmt.Println(
		fmap(
			func(x interface{}) interface{} {
				return x.(int) + 1
			},
			makeList(1),
		),
	)
	t.Error("")
}

func TestListMap3(t *testing.T) {
	fmt.Println(
		fmap(
			func(x interface{}) interface{} {
				return x.(int) + 1
			},
			makeList(1, 2, 3, 4, 5),
		),
	)
	t.Error("")
}

func TestListMap4(t *testing.T) {
	fmt.Println(
		fmap(
			func(x interface{}) interface{} {
				return x.(int) * x.(int)
			},
			makeList(1, 2, 3, 4, 5),
		),
	)
	t.Error("")
}

func TestListForEach1(t *testing.T) {
	fEach(
		func(x interface{}) {
			fmt.Println(x)
		},
		makeList(),
	)
	t.Error("")
}

func TestListForEach2(t *testing.T) {
	fEach(
		func(x interface{}) {
			fmt.Println(x)
		},
		makeList(1),
	)
	t.Error("")
}

func TestListForEach3(t *testing.T) {
	fEach(
		func(x interface{}) {
			fmt.Println(x)
		},
		makeList(1, 2, 3, 4, 5),
	)
	t.Error("")
}

func TestListForEach4(t *testing.T) {
	fEach(
		func(x interface{}) {
			fmt.Println(x.(int) * x.(int))
		},
		makeList(1, 2, 3, 4, 5),
	)
	t.Error("")
}

func TestListTree1(t *testing.T) {
	fmt.Println(newList(cons(makeList(1, 2), makeList(3, 4)())))
	t.Error("")
}

func TestListCountLeaves1(t *testing.T) {
	fmt.Println(countLeaves(makeList()))
	t.Error("")
}

func TestListCountLeaves2(t *testing.T) {
	fmt.Println(countLeaves(makeList(1, 2, 3)))
	t.Error("")
}

func TestListCountLeaves3(t *testing.T) {
	fmt.Println(countLeaves(newList(cons(makeList(1, 2), nil))))
	t.Error("")
}

func TestListCountLeaves4(t *testing.T) {
	fmt.Println(countLeaves(newList(cons(makeList(1, 2), makeList(3, 4)()))))
	t.Error("")
}

func TestListCountLeaves5(t *testing.T) {
	fmt.Println(countLeaves(makeList(1, 2, makeList(4, 5, makeList(6), makeList()), 3)))
	t.Error("")
}
