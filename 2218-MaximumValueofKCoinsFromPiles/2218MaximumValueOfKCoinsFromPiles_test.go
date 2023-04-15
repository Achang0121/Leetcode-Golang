package leetcode

import (
	"fmt"
	"testing"
)

type question2218 struct {
	args
	ans
}

type args struct {
	piles [][]int
	k     int
}

type ans struct {
	res int
}

func TestProblem2218(t *testing.T) {
	qs := []question2218{
		{
			args{[][]int{{1, 100, 3}, {7, 8, 9}}, 2},
			ans{101},
		},

		{
			args{[][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7},
			ans{706},
		},
	}
	fmt.Println("---------------------Leetcode Problem 2218 ---------------------")
	for _, q := range qs {
		_, p := q.ans, q.args
		fmt.Printf("【input】: %v \n【output】: %v\n", p, maxValueOfCoins(p.piles, p.k))
	}
	fmt.Printf("\n\n\n")
}
