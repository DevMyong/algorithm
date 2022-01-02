def solution1(n, k):
    cnt = 0
    for i in range(n):
        cnt += len(k[i + 1:]) - k[i + 1:].count(k[i])
    return cnt


def solution2(n, m, k):
    arr = [0] * 11
    for x in k:
        arr[x] = 1

    answer = 0
    for i in range(1, m + 1):
        n -= arr[i]
        answer = arr[i] * n
    return answer


if __name__ == "__main__":
    n, m = map(int, input().split())
    k = list(map(int, input().split()))
    print(solution1(n, k))
