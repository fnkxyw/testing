package main

import (
	"log"
)

type Stack struct {
	data []rune
}

func (s *Stack) Display() {
	log.Print(s.data)
}

func (s *Stack) Push(val rune) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]

}

func (s *Stack) Top() rune {
	return s.data[len(s.data)-1]
}
