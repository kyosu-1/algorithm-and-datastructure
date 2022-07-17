from collections import OrderedDict

N, X, Y, Z = map(int,input().split())

A = list(map(int,input().split()))
B = list(map(int,input().split()))
D = []
for i in range(N):
    D.append(
        {
            "id": -(i+1),
            "math": A[i],
            "en": B[i],
            "total": A[i] + B[i],
        }
    )
ans_list = []
D.sort(key=lambda x: (x["math"], x["id"]), reverse=True)

for i in range(X):
    ans_list.append(-D[i]["id"])

D = D[X:]

D.sort(key=lambda x: (x["en"], x["id"]), reverse=True)

for i in range(Y):
    ans_list.append(-D[i]["id"])

D = D[Y:]

D.sort(key=lambda x: (x["total"], x["id"]), reverse=True)

for i in range(Z):
    ans_list.append(-D[i]["id"])

ans_list.sort()
for i in ans_list:
    print(i)
