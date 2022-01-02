def solution1(a, b):
    answer = [0] * (b + 1)
    for n in range(a, b + 1):
        answer[n] += 1
        for i in range(2, int(n ** (1 / 2)) + 1):
            if n % i == 0:
                answer[i] += 1
                answer[n // i] += 1

    return answer.index(max(answer))


def solution2(a, b):
    answer = []
    for n in range(a, b + 1):
        answer.append(n)
        for i in range(2, int(n ** (1 / 2)) + 1):
            if n % i != 0:
                continue
            answer.append(i)

            if i == n // i:
                continue
            answer.append(n // i)

    max_v, max_i = 0, 0
    for i in range(2, b + 1):
        if answer.count(i) > max_v:
            max_v = answer.count(i)
            max_i = i
    return max_i


def solution3(a, b):
    if a != b:
        return 2

    for i in range(2, int(a ** (1 / 2)) + 1):
        if a % i == 0:
            return i

    return a


if __name__ == "__main__":
    a, b = map(int, input().split())
    print(solution3(a, b))
