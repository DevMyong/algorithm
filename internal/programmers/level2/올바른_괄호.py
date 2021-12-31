def solution(s):
    if len(s) % 2 == 1:
        return False

    stack = []
    for ch in s:
        if ch == "(":
            stack.append(ch)
        elif ch == ")":
            if len(stack) > 0 and stack[-1] == "(":
                stack.pop()
            else:
                return False

    return len(stack) == 0