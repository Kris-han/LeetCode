# 299. Bulls and Cows

## 题目

> You are playing the following Bulls and Cows game with your friend: You write down a number and ask your friend to guess what the number is. Each time your friend makes a guess, you provide a hint that indicates how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and how many digits match the secret number but locate in the wrong position (called "cows"). Your friend will use successive guesses and hints to eventually derive the secret number.
>
> Write a function to return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows. 
>
> Please note that both secret number and friend's guess may contain duplicate digits.
>
> Example 1:
>
> Input: secret = "1807", guess = "7810"
>
> Output: "1A3B"
>
> Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
> Example 2:
>
> Input: secret = "1123", guess = "0111"
>
> Output: "1A1B"
>
> Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.
> Note: You may assume that the secret number and your friend's guess only contain digits, and their lengths are always equal.

## 题意

在这题中，输入是两个字符串，这两个字符串长度相等，且每一个字符都是[0, 9]的数字。

题目需要我们做的是，计算两个字符串中数字相同且位置相同的数字个数，另外还要计算数字相同但位置不同的数字个数。  

## 解题思路

首先，计算两个字符串中数字相同且位置相同的数字个数是很简单的，只需要遍历一遍这两个字符串，判断是否相等就完事了。  

我们的重点在于，**如何计算数字相同但是位置不同的数字个数**。

因为这两个字符串，里面包含的数字都是[0, 9]，所以如果不考虑空间复杂度，我们可以设置两个长度为10的hash表A和B，初始值设置为0。  

在表A中，将其中一个字符串中的每一个数字都作为下标，记录下来，只要这个数字出现了，这个下标对应的值就加1，表B也同样这么做。

举个例子：

```
两个字符串分别为
1，3，2，4，2
1，4，5，2，1
```

那么对于hash表A，我们需要把下标为`1,3,2,4,1`的值都加1，其余的不变，则此时hash表A为：

```
0,2,1,1,1,0,0,0,0,0
```

对于hash表B，也是同样的操作，我们把下标为`1,4,5,2,1`的值也都加1：

```
0,2,1,0,1,1,0,0,0,0
```

此时，我们的两个hash表，就分别记录了两个字符串中出现过的数字有哪些，并且出现了多少次。

回到我们的目标，**计算数字相同但是位置不同的数字个数**。在这两个hash表中，下标对应了字符串中含有的数字，hash表的值代表了这个下标数字在原字符串中**出现了几次**。

比如对于下标为1的hash表A，`A[1] = 2`，那么就表示1这个数字，在第一个字符串中，出现了两次。

假设同样的一个数字，在字符串A中出现了3次，在字符串B中出现了2次，那么我们应该取两者的最小值，并且可以判断**这个数字在AB字符串中可能符合数字相同位置不同，也可能符合数字相同位置也相同的情况2次**。

所以我们只要把这些所有可能的结果加起来，再减去**数字相同位置也相同**的情况，就能得到我们的**数字相同但位置不同**的答案。

以上面的hash表为例，我们遍历这两个hash表，并且取`min(A[i], B[i])`，然后把所有的结果加起来，就是**这两个字符串数字相同的个数**，然后我们再减去**位置相同**的情况，就能得到**数字相同且位置不同**的答案了。

## 更进一步的解题思路

在之前的解题思路中，我们用到了两个hash表。

在这个方法中，我介绍如果用一个hash表解决问题。

对于**计算数字相同且位置也相同的数字个数**，是一样的，也需要遍历这两个字符串。

关键在于这两行代码：

```
bucket[chS[i]] += 1
bucket[chG[i]] -= 1
```

我们在这个hash表中，同样也是把字符串中的数字作为hash表的下标，然后我们把第一个字符串中出现的数字，在对应的位置加1，第二个字符串中出现的数字，在对应的位置减1。

作个假设，如果数字1在第一个字符串中出现了3次，在第二个字符串中出现了2次，那么记录到这个bucket表中，就是`bucket[1] = 3 - 2`。

此时，`bucket[1] = 1`，他的意义在于，**记录了字符串A中未能与字符串B匹配的数字个数**。

当然，`bucket[i]`的值也可能是0，这代表了这个数字`i`在字符串A、B中出现了相同的次数。

那么如果`bucket[i]`的值是负数呢，那么就代表了数字`i`在字符串B中出现的次数更多。  

但是我们可以不用考虑这些，我们只需要考虑正数的情况，也就是**数字`i`在字符串A中比字符串B多了几个**。

并且，因为`bucket[i]`的值代表了**字符串A中未能与字符串B匹配的数字个数**，那么这个值在加上**字符串A与字符串B匹配的数字个数**，将会等于数组的长度。

所以，我们可以用一个hash表解出这道题目。