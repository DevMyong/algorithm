def solution1(n):
    cnt = 0
    i = 665
    while cnt < n:
        i += 1
        if "666" in str(i):
            cnt += 1

    return i


def failed_solution(n):
    num = ""
    for i in range(n // 19 + 1):
        for j in range(n % 19):
            if j < 6:
                num = str(i) + str(j) + "666"
            elif j < 16:
                num = str(i) + "666" + str(j - 6)
            elif j < 19:
                num = str(i) + str(j - 9) + "666"
        if n % 19 == 0:
            num = str(i) + "0666"

    return num.lstrip('0')


if __name__ == "__main__":
    n = int(input())
    print(solution2(n))
