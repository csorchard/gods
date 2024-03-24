package main

import "fmt"

const ArraySize = 7

type HashMap struct {
	array [ArraySize]*LinkedList
}

func Init() *HashMap {
	hp := &HashMap{}
	for i := range hp.array {
		hp.array[i] = &LinkedList{}
	}

	return hp
}

func (h *HashMap) Insert(k string) {
	h.array[hash(k)].insert(k)
}

func (h *HashMap) Search(k string) {
	ans := h.array[hash(k)].search(k)
	fmt.Println(ans)
}

func (h *HashMap) Delete(k string) {
	h.array[hash(k)].delete(k)
}

func hash(key string) int {
	sum := 0
	for _, i := range key {
		sum += int(i)
	}

	return sum % ArraySize
}

func main() {
	hm := Init()
	hm.Insert("A")
	hm.Insert("B")
	hm.Insert("C")

	hm.Search("A")
	hm.Search("Z")
}
