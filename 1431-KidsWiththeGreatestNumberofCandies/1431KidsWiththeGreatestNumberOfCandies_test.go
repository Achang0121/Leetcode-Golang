package leetcode

import (
	"fmt"
	"testing"
)

type question1431 struct {
	args
	ans
}

type args struct {
	candies      []int
	extraCandies int
}

type ans struct {
	result []bool
}

func TestProblem1431(t *testing.T) {
	qs := []question1431{
		{
			args{[]int{2, 3, 5, 1, 3}, 3},
			ans{[]bool{true, true, true, false, true}},
		},

		{
			args{[]int{4, 2, 1, 1, 2}, 1},
			ans{[]bool{true, false, false, false, false}},
		},
	}
	fmt.Println("---------------------Leetcode Problem 1431 ---------------------")
	for _, q := range qs {
		_, p := q.ans, q.args
		fmt.Printf("【input】: %v \n【output】: %v\n", p, kidsWithCandies(p.candies, p.extraCandies))
	}
	fmt.Printf("\n\n\n")
}
