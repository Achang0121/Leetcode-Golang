# [2218. Maximum Value of K Coins From Piles](https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/)

**Hard**

## Description

There are n piles of coins on a table. Each pile consists of a positive number of coins of assorted denominations.

In one move, you can choose any coin on top of any pile, remove it, and add it to your wallet.

Given a list piles, where piles[i] is a list of integers denoting the composition of the ith pile from top to bottom, and a positive integer k, return the maximum total value of coins you can have in your wallet if you choose exactly k coins optimally.

## Example

![](https://assets.leetcode.com/uploads/2019/11/09/e1.png)

`Input: piles = [[1,100,3],[7,8,9]], k = 2`

`Output: 101`

    Explanation:
    The above diagram shows the different ways we can choose k coins.
    The maximum total we can obtain is 101.
---

## 中文大意

桌上有N堆硬币。每堆硬币由一定数量的各种面额的硬币组成。

在一次操作中，可以从任意一堆硬币的顶部选择一枚硬币，放入钱包。

给定一个硬币堆的列表和正整数k（表示取的硬币数量），`piles[i]`表示一个整型列表，元素顺序对应硬币堆从上到下；选择k个硬币，返回最大的硬币总面值。
