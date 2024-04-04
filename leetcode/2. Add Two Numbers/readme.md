# (RU) Чему я научился у задачи?

Почему-то привык спешить после 1-ой задачи, и иногда начинаю решать эту неоптимально просто по привычке :( Невнимательно читаю условие с мыслями, что знаю почти наизусть, как решаются первые несколько десятков задач с литкода, если не первые 100-200. Так что похоже, если я хочу провести время с пользой, придется останавливаться даже на таких простых задачах и внимательно и осторожно в них заново вникать.

В этой задаче помимо фишки с "неудобной" вершиной новосозданного l3 есть куда более менее заметная ловушка - можно забыть, что когда кончились l1 и l2, мог остаться остаток единица, который надо дописать в l3. Как всегда, нужно смотреть "за краями".

---

https://leetcode.com/problems/add-two-numbers/

# 2. Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



### Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

### Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
 

### Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.