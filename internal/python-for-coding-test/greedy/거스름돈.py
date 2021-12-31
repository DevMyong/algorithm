def solution1(n):
    cnt = 0
    while n > 0:
        k = 0
        if n >= 500:
            k = 500
        elif n >= 100:
            k = 100
        elif n >= 10:
            k = 10

        cnt += n // k
        n %= k
    return cnt


def solution2(n):
    coins = [500, 100, 50, 10]
    cnt = 0
    for coin in coins:
        cnt += n // coin
        n %= coin


if __name__ == "__main__":
    n = int(input())
    print(solution1(n))
