def solution(s):
    min_s = 1000
    # 단위 선택 (2 ~ n/2)
    for n in range(1, len(s) // 2 + 2):
        # 패턴 선택
        pattern = s[:n]
        answer = ""
        cnt = 1
        for i in range(n, len(s) + n, n):
            # 패턴이 반복되면 카운트 후 반복
            # 슬라이싱은 index out of range가 안생김
            if s[i:i + n] is pattern:
                cnt += 1
                continue
            # 패턴이 아닌 경우,
            answer += str(cnt) + pattern if cnt > 1 else pattern
            pattern = s[i:i + n]
            cnt = 1
        min_s = min(len(answer), min_s)

    return min_s


if __name__ == "__main__":
    print(solution("a"))
