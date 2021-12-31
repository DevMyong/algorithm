if __name__ == '__main__':
    m = [[1 for j in range(10 + 2)] for i in range(10 + 2)]

    for y in range(1, 11):
        line = list(map(int, input().split()))
        line.insert(0, 1)
        line.append(1)
        m[y] = line

    y = 2
    x = 2

    while m[y][x] != 2 and y < 10:
        m[y][x] = 9

        if m[y][x + 1] != 1:
            x += 1
        else:
            y += 1

    if m[y][x] == 2:
        m[y][x] = 9

    for i in range(1, 11):
        for j in range(1, 11):
            print(m[i][j], end=" ")
        print('')
    print('')

