package leetcode_077

import "testing"

type Path struct {
	r []int
}

func (p *Path) Pop(r int) int {
	result := p.r[len(p.r)]
	p.r = p.r[:len(p.r)-1]
	return result
}

func (p *Path) Push(r int) {
	p.r = append(p.r, r)
}

func (p *Path) Len() int {
	return len(p.r)
}

var trace = Path{}

var result = make([][]int, 0)

// 计算组合问题 回溯算法
func Test077(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	backingTrace(s, 0, 2)
}

func backingTrace(s []int, index int, depth int) {
	if trace.Len() == depth {
		result = append(result, trace.r)
		return
	}
	r :=
	for i := range s {

	}
}

// 递归模版

// 1. 确定函数入参
// 2. 确定终止逻辑
// 3. 确定单层递归逻辑
