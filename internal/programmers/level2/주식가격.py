def solution(prices):
    n = len(prices)
    time = [0] * n
    stack = []

    for i in range(0, n):
        while stack and prices[stack[-1]] > prices[i]:
            t = stack.pop()
            time[t] = i - t
        stack.append(i)

    while stack:
        t = stack.pop()
        time[t] = n-1 - t

    return time