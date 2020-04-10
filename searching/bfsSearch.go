package main

import (
	"fmt"
	"strings"
)

func main() {
	people := make(map[string][]string)
	people["artem"] = []string{"alisa", "vova", "dima"}
	people["dima"] = []string{"masha"}

	mangoSellerName := bfsSearch("artem", people)
	fmt.Println("Mango seller is " + mangoSellerName)
}

func bfsSearch(name string, people map[string][]string) string {
	deque := New()
	deque.PushSlice(people[name])
	var searched []string

	isMangoSellerNotFound := true

	for isMangoSellerNotFound {
		person := deque.Front()
		if !contains(searched, person) && isMangoSeller(person) {
			isMangoSellerNotFound = false
			return person
		} else {
			searched = append(searched, person)
			for _, val := range people[person] {
				deque.Push(val)
			}
		}
	}

	return "not found"
}

func isMangoSeller(name string) bool {
	return strings.Contains(name, "ha")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Deque struct {
	queue []string
}

func New() *Deque {
	return &Deque{
		queue: make([]string, 0),
	}
}

func (d *Deque) Push(val string) {
	d.queue = append(d.queue, val)
}

func (d *Deque) PushSlice(val []string) {
	for _, val := range val {
		d.queue = append(d.queue, val)
	}
}

func (d *Deque) Front() string {
	if len(d.queue) <= 0 {
		return ""
	}

	node := d.queue[0]
	d.queue[0] = ""
	d.queue = d.queue[1:]

	return node
}
