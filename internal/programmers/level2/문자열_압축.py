def solution(s):
    min_s = 1000
    for n in range(1, len(s) // 2 + 2):
        pattern = s[:n]
        answer = ""
        cnt = 1

        for i in range(n, len(s) + n, n):
            if s[i:i + n] == pattern:
                cnt += 1
                continue
            answer += str(cnt) + pattern if cnt > 1 else pattern
            pattern = s[i:i + n]
            cnt = 1
        min_s = min(len(answer), min_s)

    return min_s
