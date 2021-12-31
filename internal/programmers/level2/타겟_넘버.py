def solution(numbers, target):
    tree = [0]
    for num in numbers:
        tmp = []
        for node in tree:
            tmp.append(node + num)
            tmp.append(node - num)
        tree = tmp
    answer = tree.count(target)
    return answer