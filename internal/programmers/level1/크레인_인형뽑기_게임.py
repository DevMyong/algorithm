def solution(board, moves):
    answer = 0
    stack = []
    for move in moves:
        x = move-1
        for y in (range(len(board))):
            if board[y][x]:
                if len(stack) > 0 and (stack[-1] == board[y][x]):
                    stack.pop()
                    answer = answer+2
                else:
                    stack.append(board[y][x])
                board[y][x] = 0
                break
    return answer