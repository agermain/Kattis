s = int(input())
num = [int(input()) for _ in range(s)]
final = [(x//10)**(x%10) for x in num]
print(sum(final))