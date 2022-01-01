def solution1(arr, m, k):
    arr.sort(reverse=True)

    answer = 0
    i = 0
    while m > 0:
        if i % 2 == 0:
            answer += arr[0] * k
            m -= k
        else:
            answer += arr[1]
            m -= 1
        i += 1

    return answer + m * arr[1]


def solution2(arr, m, k):
    arr.sort(reverse=True)
    count = (m // (k + 1)) * k
    count += m % (k + 1)

    result = 0
    result += count * arr[0]
    result += (m - count) * arr[1]

    return result


if __name__ == "__main__":
    n, m, k = map(int, input().split())
    arr = list(map(int, input().split()))
    print(solution2(arr, m, k))
