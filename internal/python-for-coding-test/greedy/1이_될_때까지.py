def solution1(n, k):
    cnt = 0
    while n > 1:
        if n % k == 0:
            n //= k
        else:
            n -= 1
        cnt += 1
    return cnt


def solution2(n, k):
    cnt = 0
    while n >= k:
        cnt += n % k
        n -= n % k

        if n < k:
            break
        cnt += 1
        n //= k

    return cnt + n % k - 1


if __name__ == "__main__":
    n, k = map(int, input().split())
    print(solution2(n, k))
