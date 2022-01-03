from datetime import datetime


def solution(n, arr):
    for _ in range(n):
        date = input().split()
        s = datetime.strptime(date[0], "%m/%d")
        e = datetime.strptime(date[1], "%m/%d")
        arr.append([s, e])

    arr.sort(key=lambda x: x[1], reverse=True)
    arr.sort(key=lambda x: x[0])

    for i in range(1, n):
        last = arr[i - 1]
        last_e = last[1]

        now = arr[i]
        now_s, now_e = now[0], now[1]

        if now_s < last_e < now_e:
            return "No"
    return "Yes"


def solution2(n, arr):
    for i in range(n):
        date = input().split()
        s = date[0].split("/")
        e = date[1].split("/")
        d1 = int(s[0]) * 100 + int(s[1])
        d2 = int(e[0]) * 100 + int(e[1])
        arr.append((d1, d2))

    arr.sort(key=lambda x: (x[0], -x[1]))

    for i in range(1, n):
        last = arr[i - 1]
        last_e = last[1]

        now = arr[i]
        now_s, now_e = now[0], now[1]

        if now_s < last_e < now_e:
            return "No"
    return "Yes"


if __name__ == "__main__":
    n = int(input())
    arr = []
    print(solution2(n, arr))
