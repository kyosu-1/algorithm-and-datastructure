N = int(input())
sx, sy, tx, ty = map(int, input().split())
C = [tuple(map(int,input().split())) for i in range(N)]
Q = []
V = [0]*N

for i, (x, y, r) in enumerate(C):
    if (x - sx)**2 + (y - sy)**2 == r**2:
        Q.append(i)
        V[i] = 1

for q in Q:
    x,y,r = C[q]
    if (x - tx)**2 + (y - ty)**2 == r**2:
        print("Yes")
        exit()

    for j, (a, b, R) in enumerate(C):
        if V[j]:
            continue
        if (r - R) ** 2 <= (a - x)**2 + (b - y)**2 <= (r + R)**2:
            V[j] = 1
            Q.append(j)

print('No')

# TLE解法
# import dataclasses
# from typing import List, Set
# 
# 
# @dataclasses.dataclass(frozen=True)
# class Point:
#     """二次元座標を表すクラス"""
#     x: int
#     y: int
# 
#     def distance(self, point: 'Point') -> float:
#         """2点間の距離の二乗を返す"""
#         return (self.x - point.x)**2 + (self.y - point.y)**2
# 
# 
# @dataclasses.dataclass(frozen=True)
# class Circle:
#     """円を表すクラス"""
#     c: Point
#     r: int
# 
#     def is_point_in(self, point: 'Point') -> bool:
#         """円上に点があるかどうかを返す"""
#         return self.c.distance(point) == self.r**2
# 
#     # 2円が共有店を持つかどうかを返す
#     def is_intersect(self, circle: 'Circle') -> bool:
#         """2円が共有店を持つかどうかを返す"""
#         return (self.r - circle.r)**2 <= self.c.distance(circle.c) <= (self.r + circle.r)**2
# 
# 
# if __name__ == '__main__':
#     N = int(input())
#     s_x, s_y, t_x, t_y = map(int,input().split())
#     s = Point(s_x, s_y)
#     t = Point(t_x, t_y)
#     circle_collection = []
#     for _ in range(N):
#         x, y, r = map(int,input().split())
#         c = Circle(Point(x, y), r)
#         circle_collection.append(c)
# 
#     divided_circle_set_collection: List[Set[Circle]] = []
#     for c in circle_collection:
#         tmp = 0
#         for circle_set in divided_circle_set_collection:
#             for other_c in circle_set:
#                 if c.is_intersect(other_c):
#                     circle_set.add(c)
#                     tmp = 1
#                     break
#             if tmp == 1:
#                 break
#         else:
#             circle_set = set([c])
#             divided_circle_set_collection.append(circle_set)
#     # sとtはどこの円上の点か
#     s_circle = None
#     t_circle = None
#     for circle in circle_collection:
#         if circle.is_point_in(s):
#             s_circle = circle
#         if circle.is_point_in(t):
#             t_circle = circle
#     for circle_set in divided_circle_set_collection:
#         if s_circle in circle_set and t_circle in circle_set:
#             print("Yes")
#             exit()
#     print("No")