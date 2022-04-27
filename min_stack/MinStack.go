package min_stack

type Node struct {
	val  int
	min  int
	next *Node
}
type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	node := &Node{val: val, min: val, next: nil}
	if this.head != nil {
		if this.head.min < val {
			node.min = this.head.min
		}
		node.next = this.head
	}
	this.head = node
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}
