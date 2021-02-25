package trie

type TrieNode struct {
    letter byte
    term bool
    children [26]*TrieNode
}

func MakeTrie(words []string) *TrieNode {
    root := new(TrieNode)
    for _, word := range words {
        root.Insert(word)
    }
    return root
}

func (root *TrieNode) Insert(word string) {
    if len(word) == 0 {
        root.term = true
    } else {
        idx := word[0] - 'A'
        if root.children[idx] == nil {
            root.children[idx] = new(TrieNode)
        }
        root.children[idx].Insert(word[1:])
    }
}

func (root *TrieNode) Find(word string) bool {
    if root == nil {
        return false
    } else if len(word) == 0 {
        return root.term
    } else {
        idx := word[0] - 'A'
        return root.children[idx].Find(word[1:])
    }
}
