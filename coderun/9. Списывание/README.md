https://coderun.yandex.ru/problem/cheating/

---

# 9. Списывание

Во время контрольной работы профессор Флойд заметил, что некоторые студенты обмениваются записками. Сначала он хотел поставить им всем двойки, но в тот день профессор был добрым, а потому решил разделить студентов на две группы: списывающих и дающих списывать, и поставить двойки только первым.

У профессора записаны все пары студентов, обменявшихся записками. Требуется определить, сможет ли он разделить студентов на две группы так, чтобы любой обмен записками осуществлялся от студента одной группы студенту другой группы.

# Решение

Достаточно использовать классический алгоритм проверки графа на двудольность - последовательная покраска соседних вершин в два разных цвета, если при обходе встречен сосед совпадающего цвета - NO.