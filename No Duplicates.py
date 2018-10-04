n = [n for n in input().split()]
s = set()
for x in n:
    s.add(x)
if len(n) == len(s):
    print("no")
else:
    print("yes")