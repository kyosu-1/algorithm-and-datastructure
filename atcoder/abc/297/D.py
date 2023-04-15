a, b = map(int, input().split())

count = 0

a, b = max(a, b), min(a, b)

while b != 0:
    count += a // b
    a %= b
    a, b = b, a

count -= 1
print(count)
