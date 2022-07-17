a = str(input())


for i in a:
    if a.count(i) == 1:
        print(i)
        exit()
print(-1)