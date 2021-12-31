if __name__ == '__main__':
    d = [[0 for col in range(21)] for row in range(21)]

    for i in range(1, 20):
        line = list(map(int, input().split()))
        line.insert(0, 0)
        line.append(0)
        d[i] = line

    n = int(input())

    for i in range(n):
        y, x = map(int, input().split())

        for j in range(1, 20):
            if y == j:

                for k in range(1, 20):
                    if x == k:
                        continue

                    if d[j][k] == 0:
                        d[j][k] = 1
                    else:
                        d[j][k] = 0

                continue

            if d[j][x] == 0:
                d[j][x] = 1
            else:
                d[j][x] = 0

    for i in range(1, 20):
        for j in range(1, 20):
            print(d[i][j], end=' ')
        print('')

