import re

def solution(dartResult):
    bonus = {'S': 1, 'D': 2, 'T': 3}
    option = {'': 1, '*': 2, '#': -1}

    p = re.compile('(\d+)([A-Z])([*#]?)')
    c = p.findall(dartResult)

    answer = []
    for i in range(0, len(c)):
        answer.append(int(c[i][0]))
        if (c[i][2] == '*') and (i-1 >= 0):
            answer[i - 1] *= 2
        answer[i] = answer[i] ** bonus[c[i][1]] * option[c[i][2]]
    return sum(answer)