package immutable

type Trie struct {
	root *Node
}

type Node struct {
	left *Node
	right *Node
}