import string
n = int(input())
al = set(string.ascii_lowercase)
for i in range(n):
    s = input().replace(" ", "").lower()
    c = set()
    l = set()
    for j in s:
        l.add(j)
        if j in al:
            c.add(j)
    if len(c) == 26:
        print("pangram")
    else:
        print("missing", "".join(sorted(al-l)))