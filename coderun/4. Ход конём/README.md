https://coderun.yandex.ru/problem/knight-move/

# Description (Условие)

Дана прямоугольная доска N×M. В левом верхнем углу находится шахматный конь, которого необходимо переместить в правый нижний угол доски. В данной задаче конь может перемещаться на две клетки вниз и одну клетку вправо или на одну клетку вниз и две клетки вправо.

Необходимо определить, сколько существует различных маршрутов, ведущих из левого верхнего в правый нижний угол.

# Solution (Решение)

- Time - *O(N\*M)*
- Space - *O(N\*M)*

Pass through all cells of the matrix from left to right, top to bottom from (0,0). Adding to the cells into which a knight could get from the current one, the number of paths to this cell. The answer is in the last cell.

Проход по всем клеткам матрицы слева-направо, сверху-вниз от (0,0). Добавление к клеткам, в которые из текущей мог попасть конь, числа путей к данной клетке. Ответ в последней клетке.