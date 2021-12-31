def solution(n, arr1, arr2):
    answer = []

    for a, b in zip(arr1, arr2):
        map = bin(a | b)[2:]
        if len(map) < n:
            map = '0'*(n-len(map)) + map
        map = map.replace('1', '#')
        map = map.replace('0', ' ')
        answer.append(map)

    return answer