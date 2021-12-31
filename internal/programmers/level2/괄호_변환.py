def solution(p):
    # 1. 빈문자열인 경우, 반환
    if len(p) == 0 or isCorrect(p):
        return p

    # 2. u,v <- p
    u, v = split(p)

    # 3. u가 올바른지 확인
    if isCorrect(u):
        # 3.1. u + solution(v)
        return u + solution(v)

    # 4. u가 올바르지 않을 경우
    # 4.1 ~ 4.3.
    v = "(" + solution(v) + ")"

    # 4.4.
    brace_map = {"(": ")", ")": "("}
    new_u = ""
    for brace in u[1:-1]:
        new_u += brace_map[brace]

    return v + new_u


def split(p):
    # 2.1. 균형잡힌 u 확인
    u, v = "", ""
    i = 2
    while i <= len(p):
        u = p[:i]
        if u.count('(') == u.count(')'):
            break
        i += 2

    # 2.2. p의 나머지 v
    v = ""
    if len(p) > i:
        v = p[i:]

    return u, v


def isCorrect(w):
    s = []
    braces = {')': '('}
    for ch in w:
        if s and ch in braces:
            if s[-1] == braces[ch]:
                s.pop()
            else:
                return False
        else:
            s.append(ch)
    if s:
        return False

    return True