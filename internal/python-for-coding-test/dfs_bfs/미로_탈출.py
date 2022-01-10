from collections import deque


def solution(n, m, arr):
    visit = [[1] * m for _ in range(n)]

    q = deque()
    q.append([0, 0])
    visit[0][0] = 1

    dy = [1, 0, -1, 0]
    dx = [0, 1, 0, -1]

    while q:
        now_y, now_x = q.popleft()

        for i in range(4):
            next_y, next_x = now_y + dy[i], now_x + dx[i]
            if not (0 <= next_y < n) or not (0 <= next_x < m) or visit[next_y][next_x] > 1 or not arr[next_y][next_x]:
                continue

            q.append([next_y, next_x])
            visit[next_y][next_x] = visit[now_y][now_x] + 1

        if visit[n - 1][m - 1] > 1:
            return visit[n - 1][m - 1]

    return visit[n - 1][m - 1]


if __name__ == "__main__":
    n, m = map(int, input().split())
    arr = [list(map(int, input())) for _ in range(n)]
    print(solution(n, m, arr))
