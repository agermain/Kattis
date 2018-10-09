n = int(input())
l = [int(n) for n in input().split() if int(n) >= 0]
print(sum(l)/len(l))