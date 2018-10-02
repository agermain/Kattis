s = int(input())
for i in range(s):
	n = input()
	if "Simon says" in n:
		a, b = n.split("Simon says ")
		print("%s\n" % b)