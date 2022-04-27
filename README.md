# min_stack

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
 
## Examples

```
Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
 

Constraints:

-231 <= val <= 231 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 104 calls will be made to push, pop, top, and getMin.
Accepted
962,885
Submissions
1,919,653
```

## 解析

需要實作一個 MIN STACK

需要每次找最小值要在 O(1)

關鍵點在於如果是 linked List 去實作會要找最小值是 O(n)

所以可以透過多在每個結點多紀錄當下的最小值

這樣每次查詢最小值只要找結點上紀錄的最小值

然後加入結點時 只要透過 比較當下 head 的最小值就可以知道 最小值是多少


## 實作

```golang
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

```