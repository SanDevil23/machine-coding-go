package system_design

type LRUCache struct {
    q           *DoublyLinkedList;
    keyMap      map[int]*Node;
    capacity    int;
    currSize    int;
}

type Node struct{
    data int;
    prev *Node;
    next *Node;
}

type DoublyLinkedList struct {
    head *Node;
    tail *Node;
}

// method to create the head of the linked list
func (dll *DoublyLinkedList) createHead(data int){
    node:=&Node{data: data};
    dll.head = node;
    dll.tail = node;
}

// method to insert a new node at the beginning of the LL
func (dll *DoublyLinkedList) InsertNode(data int) *Node{
    node:=&Node{data: data};
    head:=dll.head
    next:=head.next
    head.next = node
    node.next = next
    node.prev = head
    next.prev = node
    return node
}

// helper to initialize the LRU Cache struct
func Constructor(capacity int) LRUCache {
    dll := DoublyLinkedList{};
    kMap := make(map[int]*Node);
    return LRUCache{
        q:         &dll,
        keyMap:    kMap,
        capacity:  capacity,
        currSize:  1,
    }
}


func (lru *LRUCache) Get(key int) int {
    node, ok := lru.keyMap[key]
    if !ok{
        return -1
    }
    return node.data
}


func (lru *LRUCache) Put(key int, val int)  {
    // fetch the dll and key map
    dll:=lru.q
    keyMap:=lru.keyMap

    // check if map contains key
    node, ok:=keyMap[key]

    // if key is present, update the node
    if ok{
        node.data = val
        return
    } else if (dll.head!=nil && lru.currSize<lru.capacity){
        node:=dll.InsertNode(val)
        keyMap[val]=node
        lru.currSize++
        return
    } else if (lru.currSize==lru.capacity){
        tail := dll.tail
        tail = tail.prev
        tail.next = nil
        node:=dll.InsertNode(val)
        keyMap[val]=node
        return
    }

    dll.createHead(val)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */