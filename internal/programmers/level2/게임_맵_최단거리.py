from collections import deque


def solution(maps):
    dxdy = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    n, m = len(maps[0]), len(maps)

    dis = [[1] * n for _ in range(m)]

    dq = deque()
    dq.append((0, 0))
    maps[0][0] = 0
    while dq:
        x, y = dq.popleft()

        for dx, dy in dxdy:
            new_x, new_y = x + dx, y + dy

            if not(0 <= new_x < n) or not(0 <= new_y < m) or maps[new_y][new_x] == 0:
                continue

            if maps[new_y][new_x] == 1:
                dq.append((new_x, new_y))
                maps[new_y][new_x] = 0
                dis[new_y][new_x] = dis[y][x] + 1

    if dis[m - 1][n - 1] == 1:
        return -1

    return dis[m - 1][n - 1]