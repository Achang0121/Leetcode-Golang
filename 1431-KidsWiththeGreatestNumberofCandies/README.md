# [1431. Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/)

**Easy**

## Description

There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.

## Example

```shell
Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true] 
Explanation: If you give all extraCandies to:
- Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
- Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
- Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
- Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
- Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
```

---

## 中文大意

n个有糖果的孩子。给定一个数组`candies`，`candies[i]`表示第i个小朋友拥有的糖果数量，给定一个整数`extraCandies`表示你拥有的额外的糖果数量。

返回一个长度为n的布尔类型的数组`result`，`result[i]`为`true`表示如果给第i个孩子所有数量的`extraCandies`，使得该孩子拥有他们之种最多的糖果，否则为`false`。

注：多个孩子可以有最大的糖果数量。
