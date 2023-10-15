package main

import (
  "fmt"
)

type Node struct {
  val string
  prev *Node
  next *Node
}

type Queue struct {
  head *Node
  tail *Node
  size int
}

type Cache struct {
  queue Queue
  hash map[string]*Node
}

const (
  QUEUE_SIZE = 5
)

func NewCache() *Cache {
  h := &Node{}
  t := &Node{}
  h.next = t
  t.prev = h
  q := Queue{
    head: h,
    tail: t,
  }
  hash := make(map[string]*Node)
  return &Cache{queue: q, hash: hash}
}

func (c *Cache) Add(newNode *Node) {
  tmp := c.queue.head.next
  c.queue.head.next = newNode
  newNode.prev = c.queue.head
  tmp.prev = newNode
  newNode.next = tmp

  c.hash[newNode.val] = newNode

  c.queue.size++
}

func (c *Cache) Delete(node *Node) *Node {
  node.prev.next = node.next
  node.next.prev = node.prev

  delete(c.hash, node.val)
  c.queue.size--

  return node
}

func (c *Cache) Check(str string) {
  if node, ok := c.hash[str]; ok {
    tmp := c.Delete(node)
    c.Add(tmp)

  } else {
    if c.queue.size == QUEUE_SIZE {
      c.Delete(c.queue.tail.prev)
    }
    c.Add(&Node{val: str})
  }
}

func (c Cache) Display() {
  cur := c.queue.head.next
  fmt.Printf("%d [", c.queue.size)
  for cur != c.queue.tail && cur != nil {
    fmt.Printf("%s <-->", cur.val)
    cur = cur.next
  }
  fmt.Println("]")
}

func main() {
  cache := NewCache()
  items := []string{"abc", "def", "hij", "klm", "nop", "qrs", "abc", "nop"}
  for _, item := range items {
    cache.Check(item)
    cache.Display()
  }
}
