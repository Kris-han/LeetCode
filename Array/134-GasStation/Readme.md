# 134. Gas Station

## 题目

> There are N gas stations along a circular route, where the amount of gas at station i is gas[i].
>
> You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.
>
> Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.
>
> Note:
>
> If there exists a solution, it is guaranteed to be unique.
> Both input arrays are non-empty and have the same length.
> Each element in the input arrays is a non-negative integer.
> Example 1:
>
> Input: 
> gas  = [1,2,3,4,5]
> cost = [3,4,5,1,2]
>
> Output: 3
>
> Explanation:
> Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
> Travel to station 4. Your tank = 4 - 1 + 5 = 8
> Travel to station 0. Your tank = 8 - 2 + 1 = 7
> Travel to station 1. Your tank = 7 - 3 + 2 = 6
> Travel to station 2. Your tank = 6 - 4 + 3 = 5
> Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
> Therefore, return 3 as the starting index.
> Example 2:
>
> Input: 
> gas  = [2,3,4]
> cost = [3,4,3]
>
> Output: -1
>
> Explanation:
> You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
> Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
> Travel to station 0. Your tank = 4 - 3 + 2 = 3
> Travel to station 1. Your tank = 3 - 3 + 3 = 3
> You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
> Therefore, you can't travel around the circuit once no matter where you start.
>
> 链接：https://leetcode-cn.com/problems/gas-station

## 题意

这题的意思是这样的，给你两个数组，分别记录了在这个加油站能加多少油，以及去下一个加油站需要消耗多少油，并且，这些加油站是按照顺序围成一个圈的。

假设此时`i = 3`，那么gas[3]表示了在3号加油站能加的油量，cost[3]表示了从3号加油站去到4号加油站需要消耗的油量。

题目要我们做的是：一辆初始油罐的车，找一个加油站`i`，从这个加油站出发，加上这个加油站有的油`gas[i]`，然后消耗`cost[i]`的油，前往加油站`i + 1`，必须得能够回到这个加油站。答案唯一，且如果找不到这样的加油站，那么返回`-1`。

## 解法

**我们完全可以用暴力解法，遍历每一个加油站的情况，但是时间复杂度为O(n^2)，不够优雅。**



我们先来构造一个函数，`f(i) = gas[i] - cost[i]`，这个函数的意义很容易可以理解，在第`i`个加油站，你能够获取到`f(i)`这个数量的油，并且前往下一个加油站。

在这里，`f(i)`的取值有这么几个意义：

- `f(i)`为正数，那么你在这个加油站加的油是多于前往下一个加油站所需要消耗的油的，也就是说你油箱中的油量变多了
- `f(i)`为负数，那么你的油量并不够前往下一个加油站，换一句话来说，经过这个加油站之后，你的油量没有变多，甚至还被白嫖了一部分
- `f(i)`等于0，那么在经过这个加油站前往下一个加油站的过程中，你油箱中的油量是没有发生变化的



因为我们的油箱初始状态是0，并且需要走完全部的加油站，所以我们把这n个加油站的`f(i)`加起来，得到的就是最后我们油箱能够剩余的油量。如果这个值大于0，那么说明这题是有解的；如果小于0，则表示不存在有这么一个加油站。

另外，我们可以再增加一个变量，用来记录这辆汽车的当前油量。

这样，每到一个加油站，就将当前油量加上这个站能加的油量再减去前往下一个加油站所需要的油量，即

`currentGas = currentGas + gas[i] - cost[i]`

如果`currentGas`是大于等于0的，就说明这辆车能够前往`i + 1`这个加油站。

那么我们从`i = 0`这个加油站往后遍历，在这期间，如果在第`i`个加油站的时候，`currentGas < 0`，那么就说明这辆汽车没办法从起点开往`i + 1`这个加油站，我们可以把起点设置为`i + 1`。

汽车没办法开往`i + 1`这个加油站很容易理解，因为当前油量已经是负数了，所以我想解释一下为什么要把起点设置为`i+ 1`，而不是0到i中的某一个点。

因为“**如果我们把起点设置为0到i中的一个点k，能够保证这辆车到达点i的时候拥有比起点设置为0更多的油量**”这个说法成立的话，就必须保证`0-k`这段路消耗了更多的油量（因为`0-i`能够获得的油量是确定的），而只有`0-k`这段路的f(x)和是负数的，才能保证`k-i`这段路的f(x)是更大的正数。

(这段话我解释的不好，我再想想怎么解释)



但是，如果f(i)和是负数了，汽车已经没有办法到达k，所以可以证明“**如果汽车能够到达点i，那么除了起点，中间的任何一个点都是没办法到达点i的**”

所以此时我们只要能找到一个起点i，并且可以保证从i到n的currentGas是正数，那么说明这个点i，是符合要求的。

但是问题来了，现在我们只能够保证汽车可以从i为起点，开到n，怎么能够保证汽车能够从n到0到i的开回去呢？

那么我们继续做个假设，假设从0到i的这段路，存在一个点k，汽车没有办法开到点k。

**我们这里用sum(a, b)表示从f(a)到f(b)的累计**

那么我们可以得出：

`sum(0, k - 1) + sum(k, i - 1) + sum(i, n) >= 0` （式子1）

(假设这题是有解的）

并且，`sum(k, i - 1)`（式子2）一定是负的，因为如果他是正的，我们的起点完全可以从k开始，而不是从i开始。

因为式子1和式子2，我们可以得出

`sum(0, k - 1) + sum(i, n) >= 0`

而这个式子也就说明了从i开往k的f(i)的和是大于0的，也就是说这辆汽车可以从点i开往点k。

但是我们又假设我们是不能够开往点k的，所以这个假设是错误的。

因此，我们可以得出结论，**如果存在一个点i，能够开往点n，那么点i就是本题的解**。