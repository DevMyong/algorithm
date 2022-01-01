from itertools import combinations


def solution1(n, coins):
    answer = [0] * sum(coins)

    for i in range(1, n):
        combs = list(combinations(coins, i))
        for com in combs:
            answer[sum(com)] = 1
    return answer.index(0, 1)


def solution2(coins):
    coins.sort()

    target = 1
    for x in coins:
        if target < x:
            break
        target += x

    return target


if __name__ == "__main__":
    n = int(input())
    coins = list(map(int, input().split()))
    print(solution2(coins))
