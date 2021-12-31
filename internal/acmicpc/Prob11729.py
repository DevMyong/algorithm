def hanoi(n, start, middle, end):
    if n == 1:
        print(f"{start} {end}")
    else:
        hanoi(n - 1, start, end, middle)
        print(f"{start} {end}")
        hanoi(n - 1, middle, start, end)
    return


if __name__ == "__main__":
    n = int(input())
    print(2 ** n - 1)
    hanoi(n, 1, 2, 3)

