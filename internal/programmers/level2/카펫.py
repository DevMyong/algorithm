def solution(brown, yellow):
    y = 1
    while y <= yellow / y:
        if (y + 2) * 2 + yellow / y * 2 == brown:
            return [int(yellow / y + 2), (y + 2)]
        y += 1

    return []