package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type node struct {
	key    string
	value  int
	height int
	right  *node
	left   *node
}

func (subTree *node) getHeight() int {
	return subTree.height
}

func (subTree *node) recalculateHeight() {
	subTree.height = int(math.Max(float64(subTree.left.height), float64(subTree.right.height))) + 1
}

func (n *node) rotateLeft() *node {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *node) rotateRight() *node {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *node) rebalanceTree() *node {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *node) add(key string, value int) *node {
	if n == nil {
		return &node{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		n.value = value
	}
	return n.rebalanceTree()
}

func (n *node) search(key string) (int, bool) {
	if n == nil {
		return 0, false
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n.value, true
	}
}

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

type hashMap struct {
	hash map[string]int
}

type AVLTree struct {
	root *node
}

func (hash hashMap) Assign(s string, x int) {
	hash.hash[s] = x
}

func (hash hashMap) Lookup(s string) (x int, exists bool) {
	val, ok := hash.hash[s]
	return val, ok
}

func (tree *AVLTree) Assign(s string, x int) {
	tree.root = tree.root.add(s, x)
}

func (tree *AVLTree) Lookup(s string) (x int, exists bool) {
	return tree.root.search(s)

}

func parseSentence(sentence string) []string {
	var (
		res  []string
		flag = false
		word = ""
	)
	for i := 0; i < len(sentence); i++ {
		if (sentence[i] == ' ' || i == len(sentence)-1) && flag {
			if i == len(sentence)-1 && sentence[i] != ' ' {
				word += string(sentence[i])
			}
			if word != "" {
				res = append(res, word)
				word = ""
			}
			flag = false
			continue
		} else if sentence[i] == ' ' && !flag {
			continue
		} else if sentence[i] != ' ' && !flag {
			if i == len(sentence)-1 {
				word += string(sentence[i])
				res = append(res, word)
				continue
			}
			flag = true
			word += string(sentence[i])
		} else if sentence[i] != ' ' && flag {
			word += string(sentence[i])
		}
	}
	return res
}

func lex(sentence string, array AssocArray) []int {
	var (
		res          []int
		arrayOfWords []string
		cnt          = 0
	)

	arrayOfWords = parseSentence(sentence)
	//fmt.Println(arrayOfWords)
	for i := 0; i < len(arrayOfWords); i++ {
		var ok bool
		key := arrayOfWords[i]
		_, ok = array.Lookup(key)
		if !ok {
			cnt++
			array.Assign(key, cnt)
		} else {
			continue
		}
	}

	for i := 0; i < len(arrayOfWords); i++ {
		var val int
		val, _ = array.Lookup(arrayOfWords[i])
		res = append(res, val)
	}

	return res
}

func main() {
	var (
		sentence string
		hash     = make(map[string]int)
	)
	hashTable := hashMap{hash}
	assocArray := AssocArray(hashTable)

	Scan := bufio.NewScanner(os.Stdin)
	Scan.Scan()
	sentence = Scan.Text()

	res := lex(sentence, assocArray)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
}
