def solution(s):
    brackets = {")": "(", "}": "{", "]": "["}
    answer = 0

    for i in range(len(s) - 1):
        s_rotated = s[i:] + s[:i]

        # 첫번째 인덱스가 close 일 경우 넘김
        if s_rotated[0] in brackets.keys():
            continue

        stack = []
        stack.append(s_rotated[0])

        for bracket in s_rotated[1:]:
            # close bracket 일 경우
            if bracket in brackets.keys():
                if stack and brackets[bracket] == stack[-1]:
                    stack.pop()
            # open bracket 일 경우
            else:
                stack.append(bracket)

        if not stack:
            answer += 1

    return answer