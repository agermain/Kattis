import math
f = input()
d = len(f)
if d < 10:
    i, j, k = 1, 1, int(f)
    while j < k:
        i += 1
        j *= i
    print(i)
else:
    x, z = 0, 0
    while x < d-1:
        z += 1
        x += math.log(z, 10)
    print(z)