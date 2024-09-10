package guia5

import (
	"github.com/untref-ayp2/data-structures/list"
)

type HashTable struct {
	n       uint32
	buckets []*list.LinkedList[int]
}

func NewHashTable(m int) *HashTable {
	buckets := make([]*list.LinkedList[int], m)
	for i := range buckets {
		buckets[i] = list.NewLinkedList[int]()
	}
	return &HashTable{
		n:       0,
		buckets: buckets,
	}
}

func (ht *HashTable) getM() int {
	return len(ht.buckets)
}

func (ht *HashTable) hashFunction(value int) int {
	return value % ht.getM()
}

func (ht *HashTable) Insert(value int) {
	hash := ht.hashFunction(value)
	list := ht.buckets[hash]
	if list.Find(value) == nil {
		list.Append(value)
		ht.n++
	}
}

func (ht *HashTable) Exists(value int) bool {
	hash := ht.hashFunction(value)
	list := ht.buckets[hash]
	return list.Find(value) != nil
}

func (ht *HashTable) Remove(value int) {
	hash := ht.hashFunction(value)
	list := ht.buckets[hash]
	if node := list.Find(value); node != nil {
		list.Remove(value)
		ht.n--
	}
}

func (ht *HashTable) GetSize() uint32 {
	return ht.n
}