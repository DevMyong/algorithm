# 정확도 O 효율 X
def solution(food_times, k):
    s = sum(food_times)
    n = len(food_times)

    sec = 0
    i = 0
    while s > 0:
        if sec == k and food_times[i % n] > 0:
            return i % n + 1

        if food_times[i % n] > 0:
            food_times[i % n] -= 1
            sec += 1
            s -= 1
        i += 1

    return -1


if __name__ == "__main__":
    print(solution([3, 1, 1, 1, 2, 4, 3], 12))
