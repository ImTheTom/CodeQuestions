package hashmaps

import (
	"fmt"
	"time"
)

// Go has inbuilt maps and can usually get away with just using that
// make(map[interface{}]interface{})
// make(map[string]int)
// This implementation is also piggy backing on it

type HashTable struct {
	Items map[int]interface{}
}

func hash(k string) int {
	h := 0
	for i := 0; i < len(k); i++ {
		h = 31*h + int(k[i])
	}
	return h
}

func (ht *HashTable) Create() {
	if ht.Items == nil {
		ht.Items = make(map[int]interface{})
	}
}

func (ht *HashTable) Put(k string, v interface{}) {
	key := hash(k)
	ht.Items[key] = v
}

func (ht *HashTable) Remove(k string) {
	key := hash(k)
	delete(ht.Items, key)
}

func (ht *HashTable) Get(k string) interface{} {
	key := hash(k)
	return ht.Items[key]
}

func (ht *HashTable) Size() int {
	return len(ht.Items)
}

func main() {
	start := time.Now()

	fmt.Println("Running main")

	ht := HashTable{}
	ht.Create()

	ht.Put("first", 1)
	ht.Put("second", 2)
	ht.Put("third", 3)
	ht.Put("fourth", 4)

	fmt.Println(ht.Get("third"))
	fmt.Println(ht.Size())

	ht.Remove("second")
	ht.Remove("first")
	fmt.Println(ht.Size())

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
