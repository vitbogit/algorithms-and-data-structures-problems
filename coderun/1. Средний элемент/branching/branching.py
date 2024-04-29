a, b, c = map(int, input().split(" "))
if a < b:
    if c >= b:
        print(b)
    elif c > a:
        print(c)
    else: 
        print(a) 
elif a == b:
    print(a)
else: # b < a
    if c >= a:
        print(a)
    elif c > b:
        print(c)
    else:
        print(b)