package main

import (
	"math/rand"
	"strconv"
)

type Tree struct {
	parent *Tree
	left   *Tree
	value  interface{}
	right  *Tree
}

type Task func(v interface{})

func NewTree(v interface{}) *Tree {
	return &Tree{
		parent: nil,
		left:   nil,
		value:  v,
		right:  nil,
	}
}

func RandomPopulate(numNode int) *Tree {
	trees := make([]*Tree, numNode+1)
	for i := 1; i <= numNode; i++ {
		trees[i] = NewTree(rand.Intn(numNode))
	}
	for i := 1; i < numNode/2; i++ {
		trees[i].appendLeft(trees[2*i])
		trees[i].appendRight(trees[2*i+1])
	}
	return trees[1]
}

func (self *Tree) appendLeft(l *Tree) {
	self.left = l
	l.parent = self
}

func (self *Tree) appendRight(r *Tree) {
	self.right = r
	r.parent = self
}

func (self *Tree) getRoot() *Tree {
	if self.parent == nil {
		return self
	} else {
		return self.parent.getRoot()
	}
}

func (self *Tree) chanTraverse(preChan, inChan, postChan chan interface{}) {
	chanTraverseImp(self, preChan, inChan, postChan)
	if preChan != nil {
		preChan <- "done"
	}
	if inChan != nil {
		inChan <- "done"
	}
	if postChan != nil {
		postChan <- "done"
	}
}

func chanTraverseImp(self *Tree, preChan, inChan, postChan chan interface{}) {
	if self != nil {
		if preChan != nil {
			preChan <- self.value
		}
		chanTraverseImp(self.left, preChan, inChan, postChan)
		if inChan != nil {
			inChan <- self.value
		}
		chanTraverseImp(self.right, preChan, inChan, postChan)
		if postChan != nil {
			postChan <- self.value
		}
	}
}

func (self *Tree) traverse(preTask, inTask, postTask Task) {
	if self == nil {
		return
	}
	if preTask != nil {
		preTask(self.value)
	}
	self.left.traverse(preTask, inTask, postTask)
	if inTask != nil {
		inTask(self.value)
	}
	self.right.traverse(preTask, inTask, postTask)
	if postTask != nil {
		postTask(self.value)
	}
}

func (self *Tree) intJoin(result chan string) {
	if self == nil {
		result <- ""
	} else {
		lChan := make(chan string)
		rChan := make(chan string)
		go self.left.intJoin(lChan)
		go self.right.intJoin(rChan)
		result <- (<-lChan + " " + strconv.Itoa(self.value.(int)) + " " + <-rChan)
	}
}
