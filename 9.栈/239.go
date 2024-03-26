package main

import (
	"fmt"
)

type myqueue struct {
	queue []int
}

func newMyqueue() *myqueue {
	return &myqueue{
		queue: []int{},
	}
}

func (p *myqueue) front() int {
	return p.queue[0]
}

func (p *myqueue) back() int {
	return p.queue[len(p.queue)-1]
}

func (p *myqueue) empty() bool {
	if len(p.queue) != 0 {
		return false
	} else {
		return true
	}
}

func (p *myqueue) pop(val int) {
	if !p.empty() && val == p.queue[0] {
		p.queue = p.queue[1:]
	}
}

func (p *myqueue) push(val int) {
	for !p.empty() && val > p.back() {
		p.queue = p.queue[:len(p.queue)-1]
	}
	p.queue = append(p.queue, val)
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	queue := newMyqueue()
	for i := 0; i < k; i++ {
		queue.push(nums[i])
	}
	res = append(res, queue.front())
	for i := k; i < len(nums); i++ {
		queue.pop(nums[i-k])
		queue.push(nums[i])
		res = append(res, queue.front())
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}
