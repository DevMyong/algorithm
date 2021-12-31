def solution(n):
    answer = fib(n) % 1234567
    return answer


def fib(n):
    a, b = 0, 1
    for i in range(n-1):
        a, b = b, a+b

    return b