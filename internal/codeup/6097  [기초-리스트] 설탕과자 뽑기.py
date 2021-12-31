if __name__ == '__main__':
    h, w = map(int, input().split())
    m = [[0 for j in range(w + 2)] for i in range(h + 2)]

    n = int(input())

    for i in range(n):
        l, d, x, y = map(int, input().split())

        for k in range(l):
            if d == 0:
                m[x][y + k] = 1
            else:
                m[x + k][y] = 1

    for i in range(1, h + 1):
        for j in range(1, w + 1):
            print(m[i][j], end=' ')
        print('')

