def solution(s):
    answer = 0
    for num in s:
        num = int(num)
        if num <= 1 or answer <= 1:
            answer += num
        else:
            answer *= num
    return answer


if __name__ == "__main__":
    s = input()
    print(solution(s))
