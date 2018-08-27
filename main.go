package main

import (
	"bytes"
	"log"
	"sort"
	"strconv"
)

type item struct {
	v    string
	w    int
	code string
}

type itemTreeNode struct {
	prev  *itemTreeNode
	value *item
	next  *itemTreeNode
}

func (m itemTreeNode) toString() string {
	return m.value.v + " " + strconv.Itoa(m.value.w) + " " + m.value.code
}

func (m itemTreeNode) isLeaf() bool {
	return m.prev == nil && m.next == nil
}
func dumpHeapValue(item itemTreeNode) {
	log.Printf("%v  %s %d %s", item, item.value.v, item.value.w, item.value.code)
}

func nSpaces(spaces int) string {
	var buffer bytes.Buffer
	for i := 0; i < spaces; i++ {
		buffer.WriteString("    ")
	}
	return buffer.String()
}

func dumpNode(item *itemTreeNode, row int, col int) {
	log.Printf("\t\t\t %s [%d] %s \n", nSpaces(col), row, item.toString())

	if item.prev != nil {
		dumpNode(item.prev, row+1, col)
	}
	if item.next != nil {
		dumpNode(item.next, row+1, col+1)
	}
}

func setCodes(item *itemTreeNode, suffix string) {
	if item.isLeaf() {
		item.value.code = suffix
	} else if item.prev != nil {
		setCodes(item.prev, suffix+"0")
	}
	if item.next != nil {
		setCodes(item.next, suffix+"1")
	}

}

func main() {
	log.Printf("encoding")
	inputValues := []item{item{"a", 5, ""}, item{"b", 9, ""}, item{"c", 12, ""}, item{"d", 13, ""}, item{"e", 16, ""}, item{"f", 45, ""}}

	var heapValues []itemTreeNode

	for i := 0; i < len(inputValues); i++ {
		heapValues = append(heapValues, itemTreeNode{nil, &inputValues[i], nil})
	}

	for len(heapValues) > 1 {
		first := heapValues[0]
		second := heapValues[1]

		parentNode := itemTreeNode{&first, &item{"", first.value.w + second.value.w, ""}, &second}
		heapValues[0] = parentNode
		for j := 2; j < len(heapValues); j++ {
			heapValues[j-1] = heapValues[j]
		}
		heapValues = heapValues[:len(heapValues)-1]

		sort.Slice(heapValues, func(i, j int) bool {
			return heapValues[i].value.w < heapValues[j].value.w
		})

		log.Printf("======================================")
		for j := 0; j < len(heapValues); j++ {
			dumpHeapValue(heapValues[j])
		}

	}

	root := heapValues[0]

	setCodes(&root, "")
	dumpNode(&root, 0, 0)

	/*
		    f          0
		    c          100
		    d          101
		    a          1100
		    b          1101
			e          111
	*/
}
