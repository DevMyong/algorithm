def solution(s):
    dy = [2, 2, 1, 1, -1, -1, -2, -2]
    dx = [-1, 1, -2, 2, -2, 2, -1, 1]

    y = int(s[1])
    x = int(chr(ord(s[0]) - 48))

    answer = 0
    for i in range(8):
        if 1 < y + dy[i] <= 8 and 1 < x + dx[i] <= 8:
            answer += 1
    return answer


if __name__ == "__main__":
    s = input()
    print(solution(s))
