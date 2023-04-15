def count_fast(N):
    divisor_counts = [0] * (N + 1)
 
    for x in range(1, N + 1):
        for multiple in range(x, N + 1, x):
            divisor_counts[multiple] += 1
 
    ans = 0
    for i in range(1, N):
        X = i
        Y = N - i
        ans += divisor_counts[X] * divisor_counts[Y]
 
    return ans
 
if __name__ == "__main__":
    N = int(input())
    result = count_fast(N)
    print(result)
