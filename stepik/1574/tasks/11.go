package tasks

import (
	"fmt"
	"strconv"
)

func main() {
	var in string
	_, _ = fmt.Scan(&in)
	fmt.Println(Task11(in))
}

func Task11(in string) string {
	chars := stack[string]()
	chars.Push("Success")
	chars.Push("Success")
	brackets := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for pos, char := range in {
		switch char {
		case '(', '[', '{':
			chars.Push(strconv.Itoa(pos + 1))
			chars.Push(string(char))
		case ')', ']', '}':
			if chars.Length() == 0 || chars.Get() != brackets[string(char)] {
				return strconv.Itoa(pos + 1)
			} else {
				chars.Pop()
				chars.Pop()
			}
		}
	}
	chars.Pop()
	return chars.Get()
}

// Служебное - имплементация стека
type stackBackend[T any] struct {
	Push   func(T)
	Pop    func() T
	Get    func() T
	Length func() int
}

func stack[T any]() stackBackend[T] {
	slice := make([]T, 0)
	return stackBackend[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Get: func() T {
			res := slice[len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
	}
}
