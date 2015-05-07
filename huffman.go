package main
 
import (
    "container/heap"
    "fmt"
)

type HuffmanTree interface {
	Freq() int
}
 
type HuffmanLeaf struct {
	freq  int
	value rune
}
 
type HuffmanNode struct {
	freq int
	left, right HuffmanTree
}
 
func (self HuffmanLeaf) Freq() int {
	return self.freq
}
 
func (self HuffmanNode) Freq() int {
	return self.freq
}
 
type treeHeap []HuffmanTree
 
func (th treeHeap) Len() int {
	return len(th)
}

func (th treeHeap) Less(i, j int) bool {
	return th[i].Freq() < th[j].Freq()
}

func (th *treeHeap) Push(ele interface{}) {
	*th = append(*th, ele.(HuffmanTree))
}

func (th *treeHeap) Pop() (po interface{}) {
	po = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return
}

func (th treeHeap) Swap(i, j int){
	 th[i], th[j] = th[j], th[i]
}

func buildTree(symFreqs map[rune]int) HuffmanTree {
	var trees treeHeap
	for c, f := range symFreqs {
		trees = append(trees, HuffmanLeaf{f, c})
	}
	heap.Init(&trees)
	for trees.Len() > 1 {
		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)
		heap.Push(&trees, HuffmanNode{a.Freq() + b.Freq(), a, b})
	}
	return heap.Pop(&trees).(HuffmanTree)
}
 
func printCodes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		fmt.Printf("%s",string(prefix))
	case HuffmanNode:
		prefix = append(prefix, '0')
	        printCodes(i.left, prefix)
	        prefix = prefix[:len(prefix)-1]
	        prefix = append(prefix, '1')
	        printCodes(i.right, prefix)
	        prefix = prefix[:len(prefix)-1]
	}
}

func printCount(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		fmt.Printf("%c\t%d\n", i.value, i.freq)
	case HuffmanNode:
		prefix = append(prefix, '0')
	        printCount(i.left, prefix)
	        prefix = prefix[:len(prefix)-1]
	        prefix = append(prefix, '1')
	        printCount(i.right, prefix)
	        prefix = prefix[:len(prefix)-1]
	}
}
 
func main(){
	stringcode := "aaabbbcddddefff"
	fmt.Println("\n\n\nHUFFMAN CODE OF STRING: ",stringcode,"\n\n")
	symFreqs := make(map[rune]int)
	for _, c := range stringcode {
		symFreqs[c]++
	}
	tree := buildTree(symFreqs)
	fmt.Println("LETTERS\tCOUNT")
	printCount(tree, []byte{})
	fmt.Println("\n\n***Encode***")
	printCodes(tree, []byte{})
	fmt.Println("\n\n\n")
}
