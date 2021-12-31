if __name__ == '__main__':
    h, b, c, s = input().split()

    size = int(h) * int(b) * int(c) * int(s) / (8 * 1024 * 1024)
    print("%.1f" % size, "MB")

