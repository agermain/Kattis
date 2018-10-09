ins = []
win = []
for x in range(int(input())):
	a, b = input().split()
	if len(ins) < 12 and a not in ins:
		s = a + " " + b
		win.append(s)
		ins.append(a)
print('\n'.join(win))
