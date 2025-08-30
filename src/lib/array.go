package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type MyArray struct {
	length   int
	capacity int
	data     []int
}

func (a *MyArray) InitWithStr(input string) bool {
	var tmp int64
	var err error
	if len(input) > CAPACITY {
		fmt.Println("Length of input is too long!!")
		return false
	}
	a.capacity = CAPACITY
	a.length = 0
	a.data = make([]int, a.length, a.capacity)
	for _, v := range strings.Split(input, ",") {
		tmp, err = strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		if err != nil {
			a.data = append(a.data, 0)
		} else {
			a.data = append(a.data, int(tmp))
		}
		a.length += 1
	}
	return true
}

func (a *MyArray) InitWithSlice(input []int) bool {
	if len(input) > CAPACITY {
		fmt.Println("Length of input is too long!!")
		return false
	}
	a.capacity = CAPACITY
	a.length = len(input)
	a.data = make([]int, a.length, a.capacity)
	for i, v := range input {
		a.data[i] = int(v)
	}
	return true
}

func (a *MyArray) GetLength() int {
	return a.length
}

func (a *MyArray) GetCapacity() int {
	return a.capacity
}

func (a *MyArray) Print() []int {
	return a.data
}

func (a *MyArray) IsEmpty() bool {
	return (a.length == 0)
}

func (a *MyArray) IsFull() bool {
	return (a.capacity == a.length)
}

func (a *MyArray) GetElem(index int) (int, bool) {
	if index >= a.length || index < 0 {
		return 0, false
	}
	return a.data[index], true
}

func (a *MyArray) SetElem(index int, value int) bool {
	if index >= a.length || index < 0 {
		return false
	}
	a.data[index] = value
	return true
}

func (a *MyArray) Extend(capacity int) bool {
	if capacity <= a.capacity {
		fmt.Println("No need to extend , capacity is ", a.capacity)
		return false
	}
	tmp := make([]int, a.length, capacity)
	copy(tmp, a.data)
	a.data = tmp
	a.capacity = capacity
	return true
}

func (a *MyArray) Insert(index int, val int) bool {
	if a.IsFull() {
		fmt.Println("Is full")
		return false
	}
	if index < 0 || index >= a.length {
		fmt.Println("Invalid index")
		return false
	}
	a.data = append(a.data, 0)
	for i := a.length - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.length += 1
	a.data[index] = val
	return true
}

func (a *MyArray) Remove(index int) bool {
	if a.IsEmpty() {
		fmt.Println("Is empty")
		return false
	}
	if index < 0 || index >= a.length {
		fmt.Println("Invalid index")
		return false
	}
	a.data[index] = 0
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length -= 1
	tmp := make([]int, a.length, a.capacity)
	copy(tmp, a.data)
	a.data = tmp
	return true
}

func (a *MyArray) SmartRemove(index int) bool {
	if a.IsEmpty() {
		fmt.Println("Is empty")
		return false
	}
	if index < 0 || index >= a.length {
		fmt.Println("Invalid index")
		return false
	}
	a.data = append(a.data[0:index], a.data[index+1:]...)
	a.length -= 1
	return true
}

func (a *MyArray) Pop() bool {
	return a.Remove(a.length - 1)
}

func (a *MyArray) Append(val int) bool {
	if a.IsFull() {
		fmt.Println("Is full")
		return false
	}
	a.data = append(a.data, val)
	a.length += 1
	return true
}
