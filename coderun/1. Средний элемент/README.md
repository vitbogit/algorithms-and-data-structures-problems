# Description (Условие)

Consider three numbers a, b and c. Let's arrange them in ascending order. Which number will stand between the other two?

Рассмотрим три числа a, b и c. Упорядочим их по возрастанию. Какое число будет стоять между двумя другими?

# Solution (Решение)

## Sort (Сортировка)

- Time - *O(1)*
- Space - *O(1)*

Create slice from the given numbers and sort it using sort.Ints().

Создаем слайс из данных трех чисел и сортируем их стандартным методом sort.Ints().

## Branching (Ветвления)

- Time - *O(1)*
- Space - *O(1)*

Consider every possible case using branching.

Рассмотрим все возможные случаи с помощью ветвлений.

## Rearrangement (Перестановки)

- Time - *O(1)*
- Space - *O(1)*

Rearrange the numbers so that a <= b <= c.

Будем переставлять данные числа, пока не будет выполнено условие a <= b <= c.