def solution(s):
    n = len(s) // 2
    left, right = int(s[:n]), int(s[n:])

    r1, r2 = 0, 0
    for i in range(n):
        r1 += left // (10 ** (n - i - 1))
        r2 += right // (10 ** (n - i - 1))

        left %= 10 ** (n - i - 1)
        right %= 10 ** (n - i - 1)
    return "LUCKY" if r1 == r2 else "READY"


def solution2(s):
    n = len(s) // 2
    return "LUCKY" if sum(map(int, s[:n])) == sum(map(int, s[n:])) else "READY"


if __name__ == "__main__":
    s = input()
    print(solution2(s))
