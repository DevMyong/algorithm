def solution(s):
    ss = sorted(list(s))

    idx = 0
    for i in range(len(ss)):
        if ss[i] >= 'A':
            idx = i
            break

    return ''.join(ss[idx:]) + str(sum(map(int, ss[:idx])))


if __name__ == "__main__":
    s = input()
    print(solution(s))
