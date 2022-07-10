import numpy as np

a, b, d = map(int,input().split())


def rotation_o(u, t, deg=False):
    if deg:
        t = np.deg2rad(t)
    R = np.array([[np.cos(t), -np.sin(t)], [np.sin(t), np.cos(t)]])
    return R.dot(u)

z = np.array([a,b])

rotationed_z = rotation_o(z, d, True)

print(*rotationed_z)