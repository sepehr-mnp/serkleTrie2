package main

import (
	"fmt"

	sep "github.com/sepehr-mnp/serkleTrie2/merkle-patricia-trie"
)

func main() {
	trie := sep.NewTrie()

	trie.Put([]byte{1, 2, 3, 4}, []byte("hello"))

	trie.Put([]byte{1, 2}, []byte("world"))
	fmt.Println("fuck")
}
