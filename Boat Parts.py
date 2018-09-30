s = input()
n = [int(x) for x in s.split()]
ts = set()
counter = 1
for i in range(n[1]):
	ts.add(input())
	if len(ts) != n[0]:
		counter += 1
if len(ts) == n[0]:
	print(counter)
else:
	print("paradox avoided")