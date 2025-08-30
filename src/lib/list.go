package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type listNode struct {
	data int
	next *listNode
}

type Mylist struct {
	length   int
	capacity int
	head     *listNode
}

func (l *Mylist) newNode(data int) *listNode {
	return &listNode{
		data: data,
		next: nil,
	}
}

func (l *Mylist) InitWithSlice(input []int) bool {
	if len(input) > CAPACITY {
		fmt.Println("input is too long")
		return false
	}
	l.capacity = CAPACITY
	l.length = len(input)
	l.head = nil

	var tmp, tail *listNode
	for _, v := range input {
		tmp = l.newNode(v)
		if tail == nil {
			l.head = tmp
		} else {
			tail.next = tmp
		}
		tail = tmp
	}
	return true
}

func (l *Mylist) InitWithStr(input string) bool {
	var data []int
	for _, v := range strings.Split(input, ",") {
		tmp, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		if err != nil {
			data = append(data, 0)
		} else {
			data = append(data, int(tmp))
		}
	}
	return l.InitWithSlice(data)
}

func (l *Mylist) Print() string {
	var out string
	tmp := l.head
	for tmp != nil {
		out += strconv.FormatInt(int64(tmp.data), 10) + ","
		tmp = tmp.next
	}
	return out
}

func (l *Mylist) GetCapacity() int {
	return l.capacity
}

func (l *Mylist) GetLength() int {
	return l.length
}

func (l *Mylist) IsEmpty() bool {
	return (l.length == 0)
}

func (l *Mylist) IsFull() bool {
	return (l.capacity == l.length)
}

func (l *Mylist) GetElem(index int) (int, bool) {
	var i int = 0
	var tmp *listNode = l.head
	if index >= l.length || index < 0 {
		return 0, false
	}
	for i != index {
		tmp = tmp.next
		i += 1
	}
	return tmp.data, true
}

func (l *Mylist) SetElem(index int, value int) bool {
	var i int = 0
	var tmp *listNode = l.head
	if index > l.length-1 || index < 0 {
		fmt.Println("Invaild index")
		return false
	}
	for i != index {
		tmp = tmp.next
		i += 1
	}
	tmp.data = value
	return true
}

func (l *Mylist) Extend(capacity int) bool {
	if capacity <= l.capacity {
		fmt.Println("No need to extend , capacity is ", l.capacity)
		return false
	}
	l.capacity = capacity
	return true
}

func (l *Mylist) Insert(index int, val int) bool {
	if l.IsFull() {
		fmt.Println("Is full")
		return false
	}
	if index > l.length-1 || index < 0 {
		fmt.Println("Invaild index")
		return false
	}
	var tmp *listNode = l.newNode(val)
	var left *listNode = l.head
	var i int = 0
	if index != 0 {
		for i != index-1 {
			left = left.next
			i = i + 1
		}
		tmp.next = left.next
		left.next = tmp
	} else {
		tmp.next = left
		l.head = tmp
	}
	l.length += 1
	return true
}

func (l *Mylist) Remove(index int) bool {
	if l.IsEmpty() {
		fmt.Println("Is empty")
		return false
	}
	if index > l.length-1 || index < 0 {
		fmt.Println("Invaild index")
		return false
	}
	var left *listNode = l.head
	var tmp *listNode = nil
	var prev *listNode = nil

	if index == 0 {
		tmp = left
		l.head = left.next
	} else {
		var i int = 0
		for i != (index - 1) {
			left = left.next
			i = i + 1
		}
		prev = left
		tmp = prev.next
		prev.next = tmp.next
	}
	tmp = nil
	l.length -= 1
	return true
}

func (l *Mylist) Append(val int) bool {
	if l.IsFull() {
		fmt.Println("Is full")
		return false
	}
	var i int = 0
	var tmp *listNode = l.head
	for i != l.length-1 {
		tmp = tmp.next
		i += 1
	}
	var right *listNode = l.newNode(val)
	tmp.next = right
	l.length += 1
	return true
}

func (l *Mylist) Pop() bool {
	return l.Remove(l.length - 1)
}
