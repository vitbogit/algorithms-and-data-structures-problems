# Description (Условие)

Each cell of an N×M rectangular table contains a number. Initially, the player is in the upper left cell. In one move, he is allowed to move to an adjacent cell either to the right or down (it is prohibited to move to the left and up). When passing through a cell, the player is taken as many kilograms of food as the number written in this cell (food is also taken for the first and last cells of his path). You need to find the minimum weight of food in kilograms, giving which the player can get into the lower right corner.

В каждой клетке прямоугольной таблицы N×M записано некоторое число. Изначально игрок находится в левой верхней клетке. За один ход ему разрешается перемещаться в соседнюю клетку либо вправо, либо вниз (влево и вверх перемещаться запрещено). При проходе через клетку с игрока берут столько килограммов еды, какое число записано в этой клетке (еду берут также за первую и последнюю клетки его пути). Требуется найти минимальный вес еды в килограммах, отдав которую игрок может попасть в правый нижний угол.

# Solution (Решение)

# Dynamic programming, store matrix (Динамическое программирование, хранить строку)

- Time - *O(N\*M)*
- Space - *O(N\*M)*

We build a dynamic programming matrix, a formula for recalculating row elements: `dp[i][j] = num + min(dp[i-1][j], dp[i][j-1])` - where `num` - the next number from the original matrix. The answer is in the last cell

Строим матрицу динамического программирования, формула для пересчета элементов строки: `dp[i][j] = num + min(dp[i-1][j], dp[i][j-1])` - где `num` - очередное число из исходной матрицы. Ответ в последней ячейке

# Dynamic programming, store row (Динамическое программирование, хранить строку)

- Time - *O(N\*M)*
- Space - *O(M)*

We store only one last constructed row of the dynamic programming matrix, which we use to construct the next row. Formula for recalculating row elements: `dp[i] = num + min(dp[i-1], dp[i])` - where `num` is the next number from the source matrix. The answer is in the last cell.

Храним только одну последнюю построенную строку матрицы динамического программирования, которую используем для построения следующей строки. Формула для пересчета элементов строки: `dp[i] = num + min(dp[i-1], dp[i])` - где `num` - очередное число из исходной матрицы. Ответ в последней ячейке.