def solution(n):
    answer = ""

    while n > 0:
        n, r = divmod(n, 3)

        if r == 0:
            n -= 1

        answer = "412"[r] + answer

    return answer