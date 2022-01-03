def solution(n, arr):
    rank = [1] * n
    for i in range(n):
        for j in range(n):
            if arr[i][0] < arr[j][0] and arr[i][1] < arr[j][1]:
                rank[i] += 1
    answer = ""
    for i in range(n):
        answer += f"{rank[i]} "

    return answer[:-1]


if __name__ == "__main__":
    n = int(input())
    arr = [list(map(int, input().split())) for i in range(n)]
    print(solution(n, arr))
