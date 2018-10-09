score = 0
winner = 0
for contestant in range (1, 6):
    grades = [int(x) for x in input().split()]
    contestant_score = sum(grades)
    if contestant_score > score:
        winner = contestant
        score = contestant_score
print(winner, score)