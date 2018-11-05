for i in range(int(input())):
    sounds = input().split()
    while True:
        s = input().split()
        if s[0] == "what":
            break
        else:
            sounds = list(filter((s[2]).__ne__, sounds))
    print(' '.join(str(x) for x in list(sounds)))