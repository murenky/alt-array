package altarray

import (
	"testing"
)

func TestAdd(t *testing.T) {
	i := [5]int{1, 2, 3, 4, 5}
	a := New(i, 6, 7)
	a.AddToStart(8)
	a.AddToEnd(9)

	j, _ := a.GetElement(6)
	if j != 6 {
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	i := [5]int{1, 2, 3, 4, 5}
	a := New(i, 6, 7)
	a.AddToStart(8)
	a.AddToEnd(9)
	e := a.SetElement(6, 12)
	if e != nil {
		panic(e)
	}
	j, _ := a.GetElement(6)
	if j != 12 {
		t.Fail()
	}
}
