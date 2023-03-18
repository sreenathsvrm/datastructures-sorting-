package main


type trieNode struct {
	childerens map[rune]*trieNode
	endOfWord  bool
}

type trie struct {
	root *trieNode
}


func (t *trie) Insert(str string) {
	// start from root
	node := t.root
	// loop through the string
	for _, char := range str {
		// if a charcter is not in the trie create the node and connect in trie
		if _, ok := node.childerens[char]; !ok {
			node.childerens[char] = &trieNode{childerens: map[rune]*trieNode{}}
		}
		// jump to that char node
		node = node.childerens[char]
	}
	// at last assign that node as end of the word
	node.endOfWord = true
}

func (t *trie) Contains(str string) bool {

	node := t.root

	// loop through the string check its corresponding node is present or not
	for _, char := range str {
		// if the char node is not found then return false
		if _, ok := node.childerens[char]; !ok {
			return false
		}

		node = node.childerens[char]
	}

	return node.endOfWord
}

// function paramter as a string to delete
func (t *trie) Remove(str string) {

	findAndDelete(str, 0, t.root)
}

// function parameter str string,index int,node *trieNode : str represent whole string , index represent each index of string , node represent parent node of each character
//  1. base condition if index is equal out of the length of string then return
//  2. initiaze vairiable with value as string's index position char and converte it into rune
//  3. check the char exist in the node if exist then get it on variable newNode and do the sub steps
//     a. call the function recursively with same string and incremented index and newNode
//     b. it's work when the function returns// check newNode is the last node by matching index with string lenght
//     c. if it's the last node then assign its endOfWord as false
//     d. and then check all newNode length is zero and its not the end of any word then delete this newNode from node
//  4. return
func findAndDelete(str string, index int, node *trieNode) {

	if index >= len(str) {
		return
	}
	char := rune(str[index])

	if newNode, ok := node.childerens[char]; ok {

		findAndDelete(str, index+1, newNode)

		if index == len(str)-1 {
			newNode.endOfWord = false
		}

		if len(newNode.childerens) == 0 && !newNode.endOfWord {
			delete(node.childerens, char)
		}
	}
}