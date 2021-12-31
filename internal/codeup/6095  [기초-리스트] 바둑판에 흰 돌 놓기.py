if __name__ == '__main__':
    n = int(input())
    d = [[0 for col in range(19)] for row in range(19)]

    for i in range(n):
        a = input().split()
        d[int(a[0]) - 1][int(a[1]) - 1] = 1

    for i in range(19):
        for j in range(19):
            print(d[i][j], end=' ')
        print('')

