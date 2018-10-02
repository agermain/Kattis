x = set()
y = set()
for i in range(3):
    p = [int(n) for n in input().split()]
    if p[0] in x:
        x.remove(p[0])
    else:
        x.add(p[0])
    if p[1] in y:
        y.remove(p[1])
    else:
        y.add(p[1])
print(x.pop(), y.pop())