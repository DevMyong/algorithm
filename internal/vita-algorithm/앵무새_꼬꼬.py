def solution(arr):
    import re
    for a in arr:
        text = re.sub('[^aeiouAEIOU]', '', a)
        if len(text) == 0:
            print("???")
        else:
            print(text)


if __name__ == "__main__":
    n = int(input())
    arr = []
    for _ in range(n):
        arr.append(input())

    solution(arr)
