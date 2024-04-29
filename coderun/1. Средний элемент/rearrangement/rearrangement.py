a, b, c = map(int, input().split(" "))
if a < b:
    a, b = b, a
# at this point, a <= b
if a < c:
    a, c = c, a
# at this point, a <= b and a <= c
if b < c:
    b, c = c, b
# at this point, a <= b <= c
print(b)