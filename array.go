package altarray

import (
	"errors"
	"reflect"
)

type Element interface{}
type Elements []Element

type AltArray struct {
	mainPart    Elements
	reversePart Elements
}

// New возвращает новый AltArray на основе списка initial
// можно использовать скалярные значения, массивы и слайсы
func New(initial ...interface{}) (result AltArray) {
	var a [0]Element
	var b [0]Element
	c := a[:]
	d := b[:]
	arr := make(Elements, 0)
	for _, i := range initial {
		it := reflect.TypeOf(i)
		switch it.Kind() {
		case reflect.Array, reflect.Slice:
			s := reflect.ValueOf(i)
			for n := 0; n < s.Len(); n++ {
				arr = append(arr, Element(s.Index(n).Interface()))
			}
		default:
			arr = append(arr, Element(i))
		}
	}

	if len(arr) > 1 {
		c = append(c, arr[0])
		for i := 1; i < len(arr); i++ {
			d = append(d, arr[i])
		}
	} else {
		d = append(d, arr[0:])
	}

	result.reversePart = c
	result.mainPart = d

	return result
}

// AddToStart добавляет новый элемент в начало массива
func (a *AltArray) AddToStart(e interface{}) {
	a.reversePart = append(a.reversePart, e)
}

// AddToEnd добавляет новый элемент в конец массива
func (a *AltArray) AddToEnd(e interface{}) {
	a.mainPart = append(a.mainPart, e)
}

// GetElement возвращает элемент массива с индексом n
// нумерация элементов начинается с 0
func (a *AltArray) GetElement(n int) (Element, error) {
	x := len(a.reversePart)
	y := len(a.mainPart)

	if n+1 > x+y || n < 0 {
		return nil, errors.New("index out or range")
	}
	if n+1 <= x {
		return a.reversePart[x-n-1], nil
	}

	return a.mainPart[n-x], nil
}

// SetElement заменяет элемент с индексом n на элемент v
// нумерация элементов начинается с 0
func (a *AltArray) SetElement(n int, v interface{}) error {
	x := len(a.reversePart)
	y := len(a.mainPart)

	if n+1 > x+y || n < 0 {
		return errors.New("index out or range")
	}
	if n+1 <= x {
		a.reversePart[x-n-1] = Element(v)
		return nil
	}

	a.mainPart[n-x] = Element(v)
	return nil
}

// GetAll возвращает новый слайс, состоящий из элементов массива a
func (a *AltArray) GetAll() Elements {
	var result [0]Element
	d := result[:]

	for i := len(a.reversePart) - 1; i >= 0; i-- {
		d = append(d, a.reversePart[i])
	}
	for i := 0; i < len(a.mainPart); i++ {
		d = append(d, a.mainPart[i])
	}

	return d
}
