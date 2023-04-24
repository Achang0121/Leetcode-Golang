# [1046. Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)

**Easy**

## Description

You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:

- If `x == y`, both stones are destroyed, and

- If `x != y`, the stone of weight `x` is destroyed, and the stone of weight `y` has new weight `y - x`.

At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.

## Example 1

```shell
Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation: 
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
```

## Example 2

```shell
Input: stones = [1]
Output: 1
```

---

## 中文大意

给定一个整型数组`stones`，`stones[i]`表示第i个石头的重量。

用这些石头做个游戏。在每一轮游戏里，我们选择两个最重的石头碰撞。假设这俩最重的石头的重量分别为`x`、`y`且`x <= y`。碰撞的结果为：

- 如果`x == y`，两块石头都被损毁

- 如果`x != y`，`x`石头损毁，`y`石头的重量剩下`y - x`。

游戏结束时，最多只剩下一块石头。

如果最后剩了一块石头，返回其重量，如果没有剩下石头，返回`0`。
