N, Q = map(int, input().split())
visited_per_day = list(map(int, input().split()))

# n日目までの累積和
visited_per_day_cumsum = [0] * (N + 1)
for i in range(N):
    visited_per_day_cumsum[i + 1] = visited_per_day_cumsum[i] + visited_per_day[i]


for _ in range(Q):
    left, right = map(int, input().split())
    print(visited_per_day_cumsum[right] - visited_per_day_cumsum[left - 1])
