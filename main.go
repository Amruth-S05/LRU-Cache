// True Least Recently Used (LRU) Cache -
// 1. If an item already exists, we need to remove it and add it to the beginning
// 2. An order of items is maintained
// 3. Deletion happens at the tail and addition happens at the head

// | Node 1 | ----------> |   Cache queue   | --------> | Node n |
// item goes in                                  afte capacity is full
//                                               last item goes out

// nodes are doubly linked list

package main

import (
  "fmt"
)

type Node struct {
  val string
  prev *Node
  next *Node
}

type Queue struct { // doubly linked list
  head *Node
  tail *Node
}

type Cache struct {
  queue Queue
  hash Hash // maintains
}

type Hash map[string]*Node

func NewCache() *Cache {
  return &Cache{
    queue: Queue{},
    hash: Hash{},
  }
}

func NewQueue() *Queue {
  head := &Node{}
  tail := &Node{}

  head.next = tail
  tail.prev = head

  return &Queue{
    head: head,
    tail: tail,
  }
}


func (c Cache) Check(str string) {
  node := &Node{}
}


func main() {
  fmt.Println("start cache")
  cache := NewCache()
  for _, word := range []string{"rat", "bat", "pete", "stew"}{
    cache.Check(word)
    cache.Display()
  }
}
