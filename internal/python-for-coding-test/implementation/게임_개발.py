def solution(n, m, y, x, now_d, arr):
    move_y = [0, 1, 0, -1]
    move_x = [-1, 0, 1, 0]
    visit = [[False] * m for _ in range(n)]
    visit[y][x] = True
    answer = 1

    while True:
        # 좌로 90' 회전
        now_d += -1 if now_d > 0 else 3
        # 보는 방향 기준으로 왼쪽으로 한칸 움직임
        dy, dx = move_y[now_d], move_x[now_d]

        # 움직일 곳이 맵 밖이면 다시 처음(90도 회전)부터 시작
        if not (0 <= y + dy < n) or not (0 <= x + dx < m):
            continue

        # 상하좌우 체크
        noway = 0
        for a, b in zip(move_y, move_x):
            if not (0 <= y + a < n) or not (0 <= x + b < m):
                continue
            if arr[y + a][x + b] == 1 or visit[y + a][x + b]:
                noway += 1
        # 더 이상 움직일 곳이 없으면 뒤로 한칸 움직임
        if noway == 4:
            a, b = move_y[now_d - 1], move_x[now_d - 1]
            if arr[y + a][x + b] != 0:
                break
            y, x = y + a, x + b
            continue

        # 육지면서 방문하지 않았던 곳이면 움직임
        if arr[y + dy][x + dx] == 0 and not visit[y + dy][x + dx]:
            visit[y + dy][x + dx] = True
            y, x = y + dy, x + dx
            answer += 1

    return answer


if __name__ == "__main__":
    n, m = map(int, input().split())
    y, x, now_d = map(int, input().split())
    arr = [list(map(int, input().split())) for _ in range(n)]

    print(solution(n, m, y, x, now_d, arr))
