package guia5

import (
	"testing"
)

func TestHashTableInsertions(t *testing.T) {
	ht := NewHashTable(11)
	values := []int{4371, 1323, 6173, 4199, 4344, 9679, 1989}
	expectedHashes := []int{4371 % 11, 1323 % 11, 6173 % 11, 4199 % 11, 4344 % 11, 9679 % 11, 1989 % 11}
	for _, val := range values {
		ht.Insert(val)
	}
	for i, val := range values {
		if !ht.Exists(val) {
			t.Errorf("Value %d expected at hash %d but not found in hash table", val, expectedHashes[i])
		}
	}
	if uint32(len(values)) != ht.GetSize() {
		t.Errorf("Expected hash table size %d, got %d", len(values), ht.GetSize())
	}
}

func TestHashTableSizeSevenInsertions(t *testing.T) {
	ht := NewHashTable(7)
	values := []int{1, 8, 27, 125, 216, 343}
	for _, value := range values {
		ht.Insert(value)
		if !ht.Exists(value) {
			t.Errorf("Value %d was inserted but does not exist in HashTable", value)
		}
	}
	if ht.GetSize() != uint32(len(values)) {
		t.Errorf("Expected HashTable size %d, got %d", len(values), ht.GetSize())
	}
	for _, value := range values {
		if !ht.Exists(value) {
			t.Errorf("Expected value %d to exist in HashTable", value)
		}
	}
}