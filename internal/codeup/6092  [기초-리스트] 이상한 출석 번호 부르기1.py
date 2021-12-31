if __name__ == '__main__':
    n = int(input())
    a = input().split()

    d = [0 for i in range(23)]
    for i in range(n):
        d[int(a[i]) - 1] += 1

    for i in range(23):
        print(d[i])

