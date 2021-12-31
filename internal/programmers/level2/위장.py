def solution(clothes):
    dic = dict()

    for name, category in clothes:
        if category in dic:
            dic[category].append(name)
        else:
            dic.setdefault(category, [name])

    combination = []
    for i in dic.values():
        combination.append(len(i) + 1)

    cnt = 1
    for i in combination:
        cnt *= i

    return cnt - 1