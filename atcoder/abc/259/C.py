S = str(input())
T = str(input())

def count_substring(s):
    a = []
    lst_str = ''
    for char in s:
        if lst_str != char:
            a.append([char, 1])
            lst_str = char
            continue
        a[len(a)-1][1] += 1
    return a


a = count_substring(S)
b = count_substring(T)

if len(a) != len(b):
    print("No")
    exit()

for i, j in zip(a, b):
    if i[0] != j[0]:
        print("No")
        exit()
    if i[1] > j[1]:
        print("No")
        exit()
    if i[1] != j[1] and i[1] == 1:
        print("No")
        exit()

print("Yes")
