def solution(n, moves):
    now_y, now_x = 1, 1
    m = {'R': (0, 1), 'L': (0, -1), 'U': (-1, 0), 'D': (1, 0)}

    for move in moves:
        dy, dx = m[move][0], m[move][1]
        if not (0 < now_y + dy < n) or not (0 < now_x + dx < n):
            continue
        now_y += dy
        now_x += dx
    return now_y, now_x


if __name__ == "__main__":
    n = int(input())
    moves = list(map(str, input().split()))
    print(solution(n, moves))
