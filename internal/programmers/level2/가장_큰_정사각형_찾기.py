def solution(board):
    n, m = len(board), len(board[0])
    answer = 0
    for y in range(n):
        for x in range(m):
            answer = max(board[y][x], answer)
            if y == 0 or x == 0:
                continue
            if board[y][x] == 1:
                board[y][x] = min(board[y - 1][x], board[y][x - 1], board[y - 1][x - 1]) + 1
                answer = max(board[y][x], answer)
    return answer ** 2