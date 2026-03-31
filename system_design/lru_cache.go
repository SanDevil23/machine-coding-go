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
func (dll *DoublyLinkedList) InsertNode(data int){
    node:=&Node{data: data};
    head:=dll.head;
    next:=head.next;
    head.next = node;
    node.next = next;
    node.prev = head;
    next.prev = node;
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


func (this *LRUCache) Get(key int) int {
    
}


func (this *LRUCache) Put(key int, value int)  {
    
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */