if __name__ == '__main__':
    w, h, b = input().split()

    size = int(w) * int(h) * int(b) / (8 * 1024 * 1024)
    print("%.2f" % size, "MB")

