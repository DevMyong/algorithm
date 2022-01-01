def solution(s):
    cnt = [1 if s[0] == '0' else 0, 1 if s[0] == '1' else 0]

    for i in range(1, len(s)):
        if s[i] == s[i - 1]:
            continue
        cnt[int(s[i])] += 1

    return min(cnt)


if __name__ == "__main__":
    s = input()
    print(solution(s))
