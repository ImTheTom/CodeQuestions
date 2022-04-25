package main

import "fmt"

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	it *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		it: iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.it.hasNext()
}

func (this *PeekingIterator) next() int {
	return this.it.next()
}

func (this *PeekingIterator) peek() int {
	return this.it.v[this.it.index]
}

func main() {
	fmt.Println("Need to run in leetcode since they implemented Iterator")
}
