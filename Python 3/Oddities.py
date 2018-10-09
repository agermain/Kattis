n = [int(input()) for _ in range(int(input()))]
for x in n:
	if x%2 == 0:
		print ("{0} is even".format(x))
	else:
		print ("{0} is odd".format(x))