def solution(n):
    answer = 0
    for h in range(n+1):
        for m in range(60):
            for s in range(60):
                if '3' in str(h) + str(m) + str(s):
                    answer += 1
    return answer


def solution2(n):
    answer = 1575 * (n + 1) + 2025 * (n//10 + 1)
    return answer


if __name__ == "__main__":
    n = int(input())
    print(solution(n))
    print(solution2(n))

