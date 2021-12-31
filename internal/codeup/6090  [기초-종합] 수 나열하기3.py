if __name__ == '__main__':
    a, m, d, n = map(int, input().split())
    last = a

    for i in range(n - 1):
        last = last * m + d
    print(last)

