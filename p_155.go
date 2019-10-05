package main

type MinStack struct {
	number []int
	top    int
	min    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, -1, int(^uint(0) >> 1)}
}

func (this *MinStack) Push(x int) {
	if x < this.min {
		this.min = x
	}
	this.top++
	if this.top >= len(this.number) {
		this.number = append(this.number, x)
	} else {
		this.number[this.top] = x
	}
}

func (this *MinStack) Pop() {
	if this.Top() == this.min {
		this.min = int(^uint(0) >> 1)
		for i := 0; i < this.top; i++ {
			if this.number[i] < this.min {
				this.min = this.number[i]
			}
		}
	}
	this.top--
}

func (this *MinStack) Top() int {
	return this.number[this.top]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	obj := Constructor()
	obj.Push(3)
}
