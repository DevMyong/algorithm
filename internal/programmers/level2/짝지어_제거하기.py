def solution(s):
    if len(s) % 2 == 1:
        return 0

    stack = []
    for i in range(len(s)):
        if stack and s[i] == stack[-1]:
            stack.pop()
        else:
            stack.append(s[i])
    if stack:
        return 0
    return 1