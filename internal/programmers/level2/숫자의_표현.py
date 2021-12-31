def solution(n):
    answer = 0

    for i in range(1, n//2 + 2):
        s = 0
        for j in range(i, n//2 + 2):
            s += j
            if s == n:
                answer += 1
                break
            elif s > n:
                break

    return answer + 1