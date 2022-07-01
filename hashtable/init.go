
//hash table
package hashtable 

import "fmt"

// Here we implement a Hash table (without using the built-in map type)

type HashTable struct {
	size int
	buckets [][]*Node
}

type Node struct {
	key string
	value string
	next *Node // this idea comes from C and makes no sense to be here. we know where the next node si
}
func (h *HashTable) Hash(key string) int {
	return int(key[0]) % h.size
}
func (h *HashTable) Print() {
	for _, bucket := range h.buckets {
		for _, node := range bucket {
			fmt.Printf("%s: %s\n", node.key, node.value)
		}
	}
}

// insert
func (h *HashTable) Insert(key string, value string) {
	var k = h.Hash(key)
	bucket := h.buckets[k]
	if bucket == nil {
		bucket = []*Node{}
	    h.buckets[k] = append(bucket, &Node{key, value, nil})
	}else{
	    // add node to the head if not empty
		h.buckets[k] = append(bucket, &Node{key,value, h.buckets[k][0]})
	}
}


// remove 
func (h *HashTable) Remove(key string) {
	bucket := h.buckets[h.Hash(key)]
	if bucket == nil {
		return
	}
	for i, node := range bucket {
		if node.key == key {
			bucket = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}
// get

func (h *HashTable) Get(key string) string {
	bucket := h.buckets[h.Hash(key)]
	if bucket == nil {
		return ""
	}
	for _, node := range bucket {
		if node.key == key {
			return node.value
		}
	}
	return ""
}
