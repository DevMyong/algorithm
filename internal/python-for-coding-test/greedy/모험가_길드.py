def solution1(n, fear):
    answer = 0

    fear.sort()
    while True:
        i = fear[0]
        if i >= len(fear):
            break
        fear = fear[i:]
        answer += 1

    return answer


def solution2(fear):
    fear.sort()
    answer = 0
    cnt = 0
    for i in fear:
        cnt += 1
        if cnt >= i:
            answer += 1
            cnt = 0

    return answer


if __name__ == "__main__":
    n = int(input())
    fear = list(map(int, input().split()))
    print(solution1(n, fear))
