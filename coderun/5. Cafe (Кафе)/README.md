# Description (Условие)

A new cafe has recently opened near Petya University, which has the following discount system: for every purchase of more than 100 rubles, the buyer receives a coupon entitling him to one free lunch (for purchases worth 100 rubles or less, the buyer does not receive such a coupon).
One day Petya came across a price list for the next N days. Having studied it carefully, he decided that he would have lunch in this cafe for all N days, and every day he would buy exactly one lunch from the cafe. However, Petya’s scholarship is small, and therefore he wants to make the most of the discount system provided so that his total costs are minimal. You need to find the minimum possible total cost of lunches and the number of days on which Petya should use coupons.

Около Петиного университета недавно открылось новое кафе, в котором действует следующая система скидок: при каждой покупке более чем на 100 рублей покупатель получает купон, дающий право на один бесплатный обед (при покупке на сумму 100 рублей и меньше такой купон покупатель не получает).
Однажды Пете на глаза попался прейскурант на ближайшие N дней. Внимательно его изучив, он решил, что будет обедать в этом кафе все N дней, причем каждый день он будет покупать в кафе ровно один обед. Однако стипендия у Пети небольшая, и поэтому он хочет по максимуму использовать предоставляемую систему скидок так, чтобы его суммарные затраты были минимальны. Требуется найти минимально возможную суммарную стоимость обедов и номера дней, в которые Пете следует воспользоваться купонами.

# Solution (Решение)

- Time - *O(N\*M)*
- Space - *O(N\*M)*

`dp[i][j]` is the minimum amount that Petya could spend in i days if there are j coupons left.

If the price of the i-th lunch <= 100, then `dp[i][j]=min(dp[i-1][j]+cost[i], dp[i-1][j+1])` , where the first option means paying for the i-th lunch, and the second option means using a coupon.

If the price of the i-th lunch > 100, then `dp[i][j]=min(dp[i-1][j-1]+cost[i], dp[i-1][j+1]) `, where the first option means paying for the i-th lunch (and receiving a coupon for it), and the second option means using a coupon.

The answer is the minimum number in the last line of dp. The response index is k1. Using backtracking, we will go back, counting the time and noting the days on which the coupons were spent.

It is worth noting that additional rows and columns are created for dp, all elements are marked with infinities, and `dp[0][0]=0`.


`dp[i][j]` - минимальная суммая, которую мог потратить Петя за i дней, если осталось j талонов.

Если цена i-ого обеда <= 100, то `dp[i][j]=min(dp[i-1][j]+cost[i], dp[i-1][j+1])`, где первый вариант означает оплату i-ого обеда, а второй вариант означает использование талона. 

Если цена i-ого обеда > 100, то `dp[i][j]=min(dp[i-1][j-1]+cost[i], dp[i-1][j+1])`, где первый вариант означает оплату i-ого обеда (с получением за это талона), а второй вариант означает использование талона. 

Ответ - минимальное число в последней строке dp. Индекс ответа - это k1. С помощью backtracking`а вернемся назад, посчитав k2 и отметив дни, в которые тратились купоны.

Стоит отметить, для dp создаются дополнительные строки и столбцы, все элементы помечаются бесконечностями, а `dp[0][0]=0`.