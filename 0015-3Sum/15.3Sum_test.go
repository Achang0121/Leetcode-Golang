package leetcode

import (
	"fmt"
	"testing"
)

type question15 struct {
	args
	ans
}

type args struct {
	nums []int
}

type ans struct {
	rst [][]int
}

func TestProblem15(t *testing.T) {
	qs := []question15{
		{
			args{[]int{0, 0, 0}},
			ans{[][]int{{0, 0, 0}}},
		},

		{
			args{[]int{-1, 0, 1, 2, -1, -4}},
			ans{[][]int{{-1, 0, 1}, {-1, -1, 2}}},
		},

		{
			args{[]int{0, 1, 1}},
			ans{[][]int{}},
		},

		{
			args{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			ans{[][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		},
	}

	fmt.Println("---------------------Leetcode Problem 15 ---------------------")
	for _, q := range qs {
		_, p := q.ans, q.args
		fmt.Printf("【input】: %v \n【output】: %v\n", p, threeSum(p.nums))
	}
	fmt.Printf("\n\n\n")
}
