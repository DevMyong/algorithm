from copy import deepcopy


def rotate(key):
    n = len(key)
    ret = [[0] * n for _ in range(n)]

    for row in range(n):
        for col in range(n):
            ret[col][n - 1 - row] = key[row][col]
    return ret


def check_suit(keying, n, m):
    for y in range(n, m - n):
        for x in range(n, m - n):
            if keying[y][x] != 1:
                return False
    return True


def insert_key(key, lock, start_y, start_x):
    new_lock = deepcopy(lock)
    for y in range(len(key)):
        for x in range(len(key)):
            new_lock[y + start_y][x + start_x] += key[y][x]
    return new_lock


def solution(key, lock):
    n, m = len(key), len(lock)
    new_lock = [[0] * (m + 2 * n) for _ in range(m + 2 * n)]
    new_m = len(new_lock)

    for y in range(n, new_m - n):
        for x in range(n, new_m - n):
            new_lock[y][x] = lock[y - n][x - n]

    for i in range(4):
        key = rotate(key)

        for y in range(1, new_m - n):
            for x in range(1, new_m - n):
                keying = insert_key(key, new_lock, y, x)
                if check_suit(keying, n, new_m):
                    return True
    return False


if __name__ == "__main__":
    print(solution([[0, 0, 0], [1, 0, 0], [0, 1, 1]], [[1, 1, 1], [1, 1, 0], [1, 0, 1]]))
