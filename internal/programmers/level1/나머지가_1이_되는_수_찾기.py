def solution(n):
    answer = n - 1

    for i in range(2, (n // 2) + 1):
        if n % i == 1:
            answer = i
            break

    return answer