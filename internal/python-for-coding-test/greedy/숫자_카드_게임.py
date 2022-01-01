def solution(n):
    num = 0

    for i in range(n):
        num = max(num, min(list(map(int, input().split()))))

    return num


if __name__ == "__main__":
    n, m = map(int, input().split())
    print(solution(n))
