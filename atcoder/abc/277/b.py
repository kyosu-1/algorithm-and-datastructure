N = int(input())
S = [str(input()) for _ in range(N)]

w1 = ["H", "D", "C", "S"]
w2 = ["A", "2" , "3" , "4" , "5" , "6" , "7" , "8" , "9" , "T" , "J" , "Q" , "K"]

if len(set(S)) != N:
    print("No")
    exit()

for i in S:
    if i[0] not in w1 or i[1] not in w2:
        print("No")
        exit()

print("Yes")