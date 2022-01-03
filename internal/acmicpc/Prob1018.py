def solution(n, m, arr):
    answer = [[32] * m for i in range(n)]

    for y in range(n - 7):
        for x in range(m - 7):
            diff_wb = 0
            diff_bw = 0
            for dy in range(8):
                now_y = y + dy
                for dx in range(8):
                    now_x = x + dx
                    if arr[now_y][now_x] != map_wb[dy][dx]:
                        diff_wb += 1
                    elif arr[now_y][now_x] != map_bw[dy][dx]:
                        diff_bw += 1
            answer[y][x] = min(diff_wb, diff_bw)
    result = 32
    for i in range(n):
        result = min(min(answer[i]), result)
    return result


if __name__ == "__main__":
    n, m = map(int, input().split())
    arr = [input() for _ in range(n)]
    map_wb = ["WBWBWBWB" if i % 2 == 0 else "BWBWBWBW" for i in range(8)]
    map_bw = ["BWBWBWBW" if i % 2 == 0 else "WBWBWBWB" for i in range(8)]
    print(solution(n, m, arr))
