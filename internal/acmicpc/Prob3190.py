from collections import deque


def solution(n, turns, board):
    direction = 0
    y, x = 0, 0
    sec = 1
    snake = deque([(y, x)])

    while sec <= 1000000:
        y += dy[direction]
        x += dx[direction]

        if not (0 <= y < n) or not (0 <= x < n) or board[y][x] == 1:
            return sec

        if board[y][x] == 0:
            tail_y, tail_x = snake.popleft()
            board[tail_y][tail_x] = 0

        board[y][x] = 1
        snake.append((y, x))

        if sec in turns:
            c = turns[sec]
            direction = (direction + d[c]) % 4

        sec += 1
    return sec


if __name__ == "__main__":
    N = int(input())
    K = int(input())

    board = [[0] * N for _ in range(N)]
    apples = [list(map(int, input().split())) for _ in range(K)]
    for apple in apples:
        board[apple[0] - 1][apple[1] - 1] = 2

    L = int(input())
    turns = {}
    for _ in range(L):
        X, C = input().split()
        turns[int(X)] = C

    d = {"L": -1, "D": 1}
    dy = [0, 1, 0, -1]
    dx = [1, 0, -1, 0]

    print(solution(N, turns, board))
