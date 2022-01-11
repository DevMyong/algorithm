from collections import deque
import sys


def solution(N, M, K, X, graph):
    visit = [0] * (N + 1)

    answer = []
    q = deque([X])
    visit[q[0]] = 1
    while q:
        node = q.popleft()

        if visit[node] > K + 1:
            break

        for g in graph[node]:
            if visit[g]:
                continue
            q.append(g)
            visit[g] = visit[node] + 1
            if visit[g] == K + 1:
                answer.append(g)

    if not answer:
        answer.append(-1)

    return answer


if __name__ == "__main__":
    N, M, K, X = map(int, sys.stdin.readline().split())
    graph = [[] for _ in range(N + 1)]
    for i in range(M):
        a, b = map(int, sys.stdin.readline().split())
        graph[a].append(b)

    res = solution(N, M, K, X, graph)

    for i in res:
        print(i)
