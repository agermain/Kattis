n = int(input())
f = []
l = []
for i in range(n):
	tmp = input().split()
	f.append((float(tmp[1]), tmp[0]))
	l.append((float(tmp[2]), tmp[0]))

f = sorted(f)[:4]
l = sorted(l)[:4]
attempts = []

for t, p in f:
	people = [p]
	i = 0
	while len(people) != 4:
		if l[i][1] != p:
			people.append(l[i][1])
			t += l[i][0]
		i += 1
	attempts.append((t, people))

attempts = sorted(attempts)
print(attempts[0][0])
for p in attempts[0][1]:
	print(p)