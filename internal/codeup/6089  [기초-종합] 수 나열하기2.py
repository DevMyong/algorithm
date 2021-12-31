if __name__ == '__main__':
    a, r, n = map(int, input().split())

    last = a * r ** (n - 1) 
    print(last)

