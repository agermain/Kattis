b = False
while True:
	n = int(input())
	if n == 0:
		break
	if b:
		print("")
	else:
		b = True
	z = [[m for m in input()] for _ in range(n)]
	l = max(len(r) for r in z)
	for i in range(n):
		z[i] += [' ']*(l-len(z[i]))
	for i in range(len(z[0])):
		rr = ""
		for j in range(len(z)-1,-1,-1):
			rr += '|' if z[j][i] == '-' else '-' if z[j][i] == '|' else z[j][i]
		print(rr.rstrip())