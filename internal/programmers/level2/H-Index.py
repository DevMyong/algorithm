def solution(citations):
    citations.sort()

    for i in range(0, len(citations)):
        h = len(citations) - i
        if citations[i] >= h:
            return h
    return 0