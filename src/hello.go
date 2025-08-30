package main

import (
	"fmt"

	"jxb_go/lib"
)

func main() {
	var list lib.Mylist
	list.InitWithStr("1, 2, 3, 4, 5")
	fmt.Println("CAP = ", list.GetCapacity())
	fmt.Println("LEN = ", list.GetLength())
	fmt.Println(list.Print())
	list.Extend(200)
	list.Insert(1, 8)
	fmt.Println(list.Print())
	list.Insert(list.GetLength()-1, 9)
	fmt.Println(list.Print())
	list.Insert(0, 111)
	fmt.Println(list.Print())
	list.Remove(0)
	fmt.Println(list.Print())
	list.Remove(3)
	fmt.Println(list.Print())
	list.Remove(list.GetLength() - 1)
	fmt.Println(list.Print())
	list.Append(12)
	fmt.Println(list.Print())
	list.Pop()
	fmt.Println(list.Print())
}
