package leetcode

import (
	"fmt"
	"testing"
)

type question1046 struct {
	args
	ans
}

type args struct {
	stones []int
}

type ans struct {
	result int
}

func TestProblem1046(t *testing.T) {
	qs := []question1046{
		{
			args{[]int{2, 7, 4, 1, 8, 1}},
			ans{1},
		},

		{
			args{[]int{1}},
			ans{1},
		},
	}
	fmt.Println("---------------------Leetcode Problem 1046 ---------------------")
	for _, q := range qs {
		_, p := q.ans, q.args
		fmt.Printf("【input】: %v \n【output】: %v\n", p, lastStoneWeight(p.stones))
	}
	fmt.Printf("\n\n\n")
}
