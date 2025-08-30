package lib

const (
	CAPACITY = 100
	LENGTH   = 10
)

type LinerListOperation interface {
	GetLength() int
	GetCapacity() int
	Extend(int) bool
	Insert(int, int) bool
	Remove(int) bool
	Append(int) bool
	Pop(int) bool
	GetElem(int) (bool bool)
	SetElem(int, int) bool
	Print() bool
	IsFull() bool
	IsEmpty() bool
	InitWithStr(string) bool
	InitWithSlice([]int) bool
}
